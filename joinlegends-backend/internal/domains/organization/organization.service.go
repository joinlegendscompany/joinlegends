package organization

import (
	"fmt"
	member_repository "go-backend-stream/internal/infrastructure/repositories/member"
	organization_repository "go-backend-stream/internal/infrastructure/repositories/organization"
	user_repository "go-backend-stream/internal/infrastructure/repositories/user"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/google/uuid"
)

type OrganizationService interface {
	CreateOrganization(props CreateOrganizationServiceProps) (*models.Organization, error)
	GetOrganization(id string) (*models.Organization, error)
	GetUserOrganizations(ownerID string) ([]models.Organization, error)
	UpdateOrganization(id string, ownerID string, name string) error
	DeleteOrganization(id string, ownerID string) error
	AddMember(orgID string, ownerID string, userID string, role string) (*models.Member, error)
	RemoveMember(orgID string, ownerID string, memberID string) error
}

type organizationService struct {
	orgRepo    organization_repository.OrganizationRepository
	memberRepo member_repository.MemberRepository
	userRepo   user_repository.UserRepository
}

func NewOrganizationService(
	orgRepo organization_repository.OrganizationRepository,
	memberRepo member_repository.MemberRepository,
	userRepo user_repository.UserRepository,
) OrganizationService {
	return &organizationService{
		orgRepo:    orgRepo,
		memberRepo: memberRepo,
		userRepo:   userRepo,
	}
}

func (s *organizationService) CreateOrganization(props CreateOrganizationServiceProps) (*models.Organization, error) {
	var user models.User
	if err := s.userRepo.GetById(props.OwnerID, &user); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("user with id %s not found", props.OwnerID)
		}
		return nil, err
	}

	org := models.Organization{
		ID:        uuid.New().String(),
		Name:      props.Name,
		OwnerID:   props.OwnerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := s.orgRepo.Create(&org)
	if err != nil {
		return nil, err
	}

	ownerMember := models.Member{
		ID:             uuid.New().String(),
		UserID:         props.OwnerID,
		OrganizationID: created.ID,
		Role:           "OWNER",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if _, err := s.memberRepo.Create(&ownerMember); err != nil {
		return nil, err
	}

	created.Members = []models.Member{ownerMember}
	return &created, nil
}

func (s *organizationService) GetOrganization(id string) (*models.Organization, error) {
	var org models.Organization
	if err := s.orgRepo.GetByID(id, &org); err != nil {
		return nil, err
	}

	var members []models.Member
	if err := s.memberRepo.GetByOrganizationID(id, &members); err != nil {
		return nil, err
	}
	org.Members = members

	return &org, nil
}

func (s *organizationService) GetUserOrganizations(ownerID string) ([]models.Organization, error) {
	var orgs []models.Organization
	if err := s.orgRepo.GetByOwnerID(ownerID, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}

func (s *organizationService) UpdateOrganization(id string, ownerID string, name string) error {
	var org models.Organization
	if err := s.orgRepo.GetByID(id, &org); err != nil {
		return err
	}

	if org.OwnerID != ownerID {
		return fmt.Errorf("only the owner can update the organization")
	}

	return s.orgRepo.UpdateName(id, name)
}

func (s *organizationService) DeleteOrganization(id string, ownerID string) error {
	var org models.Organization
	if err := s.orgRepo.GetByID(id, &org); err != nil {
		return err
	}

	if org.OwnerID != ownerID {
		return fmt.Errorf("only the owner can delete the organization")
	}

	return s.orgRepo.SoftDelete(id)
}

func (s *organizationService) AddMember(orgID string, ownerID string, userID string, role string) (*models.Member, error) {
	var org models.Organization
	if err := s.orgRepo.GetByID(orgID, &org); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("organization with id %s not found", orgID)
		}
		return nil, err
	}

	if org.OwnerID != ownerID {
		return nil, fmt.Errorf("only the owner can add members")
	}

	var user models.User
	if err := s.userRepo.GetById(userID, &user); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("user with id %s not found", userID)
		}
		return nil, err
	}

	member := models.Member{
		ID:             uuid.New().String(),
		UserID:         userID,
		OrganizationID: orgID,
		Role:           role,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	created, err := s.memberRepo.Create(&member)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *organizationService) RemoveMember(orgID string, ownerID string, memberID string) error {
	var org models.Organization
	if err := s.orgRepo.GetByID(orgID, &org); err != nil {
		if err == goe.ErrNotFound {
			return fmt.Errorf("organization with id %s not found", orgID)
		}
		return err
	}

	if org.OwnerID != ownerID {
		return fmt.Errorf("only the owner can remove members")
	}

	var member models.Member
	if err := s.memberRepo.GetByID(memberID, &member); err != nil {
		if err == goe.ErrNotFound {
			return fmt.Errorf("member with id %s not found", memberID)
		}
		return err
	}

	if member.OrganizationID != orgID {
		return fmt.Errorf("member does not belong to this organization")
	}

	if member.Role == "OWNER" {
		return fmt.Errorf("cannot remove the owner from the organization")
	}

	return s.memberRepo.DeleteByID(memberID)
}
