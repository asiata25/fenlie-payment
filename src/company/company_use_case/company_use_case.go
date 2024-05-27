package companyUseCase

import (
	"finpro-fenlie/helper"
	companyDTO "finpro-fenlie/model/dto/company"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/pkg/email"
	"finpro-fenlie/src/company"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type companyUseCase struct {
	repository company.CompanyRepository
}

// GetAll implements company.CompanyUseCase.
func (c *companyUseCase) GetAll() ([]*companyDTO.CompanyResponse, error) {
	var companies []*companyDTO.CompanyResponse

	results, err := c.repository.FindAll()
	if err != nil {
		return companies, err
	}

	for _, result := range results {
		companies = append(companies, helper.ToCompanyResponse(*result))
	}

	return companies, nil
}

// Create implements company.CompanyUseCase.
func (c *companyUseCase) Create(request companyDTO.CompanyCreateRequest) error {
	admin := entity.User{
		Name:     request.User.Name,
		Email:    request.User.Email,
		Password: request.User.Password,
		Role:     "ADMIN",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.Password = string(hashedPassword)

	company := entity.Company{
		Name:      request.Name,
		SecretKey: request.SecretKey,
		Users: []entity.User{
			admin,
		},
	}

	hashedSecret, err := bcrypt.GenerateFromPassword([]byte(company.SecretKey), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	company.SecretKey = string(hashedSecret)

	id, error := c.repository.Save(company)
	bodyMail := fmt.Sprintf(
		`<table>
    <tr>
        <td colspan="2"><h3>Company</h3></td>
    </tr>
    <tr>
        <td>ID:</td>
        <td>%s</td>
    </tr>
    <tr>
        <td>Name:</td>
        <td>%s</td>
    </tr>
    <tr>
        <td>Secret Key:</td>
        <td>%s</td>
    </tr>
    <tr>
        <td colspan="2"><h3>User</h3></td>
    </tr>
    <tr>
        <td>Name:</td>
        <td>%s</td>
    </tr>
    <tr>
        <td>Email:</td>
        <td>%s</td>
    </tr>
    <tr>
        <td>Password:</td>
        <td>%s</td>
    </tr>
</table>
	`, id, request.Name, request.SecretKey, request.User.Name, request.User.Email, request.User.Password)

	if error != nil {
		email.Send(request.User.Email, "Company Account", bodyMail)
	}
	return err
}

// Delete implements company.CompanyUseCase.
func (c *companyUseCase) Delete(id string) error {
	err := c.repository.Delete(id)
	return err
}

// GetById implements company.CompanyUseCase.
func (c *companyUseCase) GetById(id string) (*companyDTO.CompanyResponse, error) {
	company, err := c.repository.RetrieveByID(id)
	if err != nil {
		return &companyDTO.CompanyResponse{}, err
	}

	return helper.ToCompanyResponse(*company), nil
}

// Update implements company.CompanyUseCase.
func (c *companyUseCase) Update(request companyDTO.CompanyUpdateRequest) error {
	companyExisting, err := c.repository.RetrieveByID(request.ID)
	if err != nil {
		return err
	}

	if request.SecretKey == "" {
		request.SecretKey = companyExisting.SecretKey
	} else {
		hashedSecret, err := bcrypt.GenerateFromPassword([]byte(request.SecretKey), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		request.SecretKey = string(hashedSecret)
	}

	company := entity.Company{
		ID:        request.ID,
		Name:      request.Name,
		SecretKey: request.SecretKey,
	}

	err = c.repository.Update(company)
	if err = helper.CheckErrNotFound(err); err != nil {
		return err
	}

	return nil
}

func NewCompanyUseCase(repository company.CompanyRepository) company.CompanyUseCase {
	return &companyUseCase{repository}
}
