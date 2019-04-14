package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/find-job-server-golang/config"
	entities "github.com/find-job-server-golang/entites"
)

type JobService struct {
}

func (jobService *JobService) CreateJob(job entities.Job) entities.Job {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Job{})
	u1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		panic(err)
	}
	job = entities.Job{JobId: u1, Name: job.Name,Decs:job.Decs,Require:job.Require,IsActive:true }
	db.Create(&job)
	return job
}

func (jobService *JobService) FindJobWithID(ID uuid.UUID) (entities.Job,bool) {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Job{})
	var job entities.Job
	isNotFound := db.Where("job_id = ?", ID).First(&job).RecordNotFound()
	return job, isNotFound
}


func (jobService*JobService) UpdateJobWithID(ID uuid.UUID, job entities.Job) entities.Job{
	db:= config.GetPostgersDB()
	db.AutoMigrate(&entities.Job{})
	var result entities.Job
	db.Model(&result).Where("job_id = ?",ID).Updates(job)
	return result
}

func (jobService *JobService) FindJobWithName(name string) (entities.Job, bool) {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Job{})
	var job entities.Job
	isNotFound := db.Where("name = ?", name).First(&job).RecordNotFound()
	return job, isNotFound
}

// func (JobService *JobService) RemoveJobWithID(ID uuid.UUID) (entities.Job, bool) {
// 	db := config.GetPostgersDB()
// 	db.AutoMigrate(&entities.Job{})
// 	var Job entities.Job
// 	isNotFound := db.Where("Job_id = ?", ID).Update("is_active")
// 	return Job, isNotFound
// }


func (jobService *JobService) Jobs(limit ,page int) []entities.Job{
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Job{})
	var jobs []entities.Job
	db.Offset(limit*page).Find(&jobs).Limit(limit)
	return jobs
}

func (jobService *JobService) JobsCount() int {
	db := config.GetPostgersDB()
	var count int
	db.Table("Jobs").Count(&count)
	fmt.Println(count)
	return count
}
