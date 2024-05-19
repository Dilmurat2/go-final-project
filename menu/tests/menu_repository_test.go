package menu_repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"menu/internal/config"
	"menu/internal/models"
	"menu/internal/ports"
	"menu/internal/repositories"
)

type MenuRepositorySuite struct {
	suite.Suite
	Repo           ports.MenuRepository
	MongoContainer testcontainers.Container
}

func (s *MenuRepositorySuite) SetupSuite() {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "root",
			"MONGO_INITDB_ROOT_PASSWORD": "example",
		},
	}
	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	s.Require().NoError(err)
	s.MongoContainer = mongoContainer

	ip, err := mongoContainer.Host(ctx)
	s.Require().NoError(err)
	port, err := mongoContainer.MappedPort(ctx, "27017")
	s.Require().NoError(err)

	cfg := &config.Config{
		Database: config.Database{
			Host:     ip,
			Port:     port.Int(),
			User:     "root",
			Password: "example",
		},
	}

	repo, err := repositories.NewMenuRepository(cfg)
	if err != nil {
		s.T().Fatalf("failed to create repository: %v", err)
	}
	s.Repo = repo
}

func (s *MenuRepositorySuite) TearDownSuite() {
	ctx := context.Background()
	if s.MongoContainer != nil {
		if err := s.MongoContainer.Terminate(ctx); err != nil {
			s.T().Errorf("error terminating MongoDB container: %v", err)
		}
	}
}

func (s *MenuRepositorySuite) TestCreate() {
	ctx := context.Background()

	menu := &models.Menu{
		ID:          primitive.NewObjectID().Hex(),
		Name:        "Test Menu",
		Description: "Test Description",
		Items:       make([]models.Item, 0),
		IsActive:    true,
	}
	createdMenu, err := s.Repo.Create(ctx, menu)
	s.Require().NoError(err)
	s.Require().NotEmpty(createdMenu.ID)
}

func (s *MenuRepositorySuite) TestUpdate() {
	ctx := context.Background()
	menu := &models.Menu{
		ID:          primitive.NewObjectID().Hex(),
		Name:        "Test Menu",
		Description: "Test Description",
		Items:       make([]models.Item, 0),
		IsActive:    true,
	}
	createdMenu, err := s.Repo.Create(ctx, menu)
	s.Require().NoError(err)
	s.Require().NotEmpty(createdMenu.ID)

	createdMenu.Name = "Updated Test Menu"
	updatedMenu, err := s.Repo.Update(ctx, createdMenu)
	s.Require().NoError(err)
	s.Require().Equal(updatedMenu.Name, "Updated Test Menu")
}

func (s *MenuRepositorySuite) TestDelete() {
	ctx := context.Background()
	menu := &models.Menu{
		ID:          primitive.NewObjectID().Hex(),
		Name:        "Test Menu",
		Description: "Test Description",
		Items:       make([]models.Item, 0),
		IsActive:    true,
	}
	createdMenu, err := s.Repo.Create(ctx, menu)
	s.Require().NoError(err)
	s.Require().NotEmpty(createdMenu.ID)

	err = s.Repo.Delete(ctx, createdMenu.ID)
	s.Require().NoError(err)
}

func (s *MenuRepositorySuite) TestGetByID() {
	ctx := context.Background()
	menu := &models.Menu{
		ID:          primitive.NewObjectID().Hex(),
		Name:        "Test Menu",
		Description: "Test Description",
		Items:       make([]models.Item, 0),
		IsActive:    true,
	}
	createdMenu, err := s.Repo.Create(ctx, menu)
	s.Require().NoError(err)
	s.Require().NotEmpty(createdMenu.ID)

	fetchedMenu, err := s.Repo.GetByID(ctx, createdMenu.ID)
	s.Require().NoError(err)
	s.Require().Equal(fetchedMenu.ID, createdMenu.ID)
}

func (s *MenuRepositorySuite) TestGetAll() {
	ctx := context.Background()
	menu := &models.Menu{
		ID:          primitive.NewObjectID().Hex(),
		Name:        "Test Menu",
		Description: "Test Description",
		Items:       make([]models.Item, 0),
		IsActive:    true,
	}
	_, err := s.Repo.Create(ctx, menu)
	s.Require().NoError(err)

	menus, err := s.Repo.GetAll(ctx)
	s.Require().NoError(err)
	s.Require().True(len(menus) > 0)
}

func TestMenuRepository(t *testing.T) {
	suite.Run(t, new(MenuRepositorySuite))
}
