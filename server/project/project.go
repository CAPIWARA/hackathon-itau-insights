package project

import (
	"hackathon-itau-insights/server/bd"

	"github.com/mitchellh/mapstructure"
)

type Project struct {
	Id        string `mapstructure:"_id" json:"id" bson:"_id"`
	CreatedBy string `json:"createdBy"`
}

func CreateProject(userId string) (Project, error) {
	m := bd.MongoBuilder("project")
	project := Project{
		CreatedBy: userId,
	}
	id, err := m.Save(project)
	if err != nil {
		return Project{}, err
	}
	project.Id = id
	return project, nil
}

func FindById(projectId string) (Project, error) {
	var project Project

	m := bd.MongoBuilder("project")
	res, err := m.FindById(projectId)
	if err != nil {
		return Project{}, err
	}

	if err := mapstructure.Decode(res, &project); err != nil {
		return Project{}, err
	}

	return project, nil
}
