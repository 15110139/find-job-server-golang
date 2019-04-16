package jobcontrolers

import (
	"fmt"
	"strconv"
	"github.com/satori/go.uuid"
	"github.com/find-job-server-golang/service"
	"github.com/gin-gonic/gin"
	entities "github.com/find-job-server-golang/entites"
	constant "github.com/find-job-server-golang/util/constant"
	response "github.com/find-job-server-golang/util/response"

)



type JobControllers struct {

}

func (jobControllers *JobControllers) CreateJob(c *gin.Context){
	var jobObj entities.Job
	c.ShouldBindJSON(&jobObj)
	jobService := service.JobService{}
	_, isNotFound := jobService.FindJobWithName(jobObj.Name)
	if !isNotFound{
		response.RespondWithError(c,constant.JOB_ALREADY_EXISTS,500)
		return
	}else{
		newJob:= jobService.CreateJob(jobObj)
		response.RespondSuccess(c,gin.H{
			"Job":newJob,
		},200)
	}

}


func (JobControllers *JobControllers) UpdateJob(c *gin.Context){
	var dataJobUpdate entities.Job
	c.ShouldBindJSON(&dataJobUpdate)
	jobService := service.JobService{}
	_, isNotFound := jobService.FindJobWithID(dataJobUpdate.JobId)
	if isNotFound{
		response.RespondWithError(c,constant.JOB_NOT_FOUND,500)
		return
	}else{
		result:= jobService.UpdateJobWithID(dataJobUpdate.JobId,dataJobUpdate)
		response.RespondSuccess(c,gin.H{
			"Job":result,
		},200)
	}

}

// func (JobControllers *JobControllers) RemoveJob(c *gin.Context){

// 	var JobId uuid.UUID
// 	c.ShouldBindJSON(&JobId)
// 	JobService := service.JobService{}
// 	_, isNotFound := JobService.FindJobWithID(JobId)
// 	if isNotFound{
// 		response.RespondWithError(c,constant.Job_NOT_FOUND,500)
// 		return
// 	}else{
// 		result:= JobService.Re(dataJobUpdate.JobId,dataJobUpdate)
// 		response.RespondSuccess(c,gin.H{
// 			"Job":result,
// 		},200)
// 	}

// }


func (jobControllers *JobControllers) Jobs(c *gin.Context){
	jobService := service.JobService{}
	page,_ :=c.GetQuery("page")
	limit,_:=c.GetQuery("limit")
	pageInt, err1 := strconv.Atoi(page)
	if err1 != nil {
		response.RespondWithError(c, constant.PAGE_MUST_BE_NUMBER, 500)
		return
	}
	limitInt, err2 := strconv.Atoi(limit)
	if err2 != nil {
		response.RespondWithError(c, constant.LIMIT_MUST_BE_NUMBER, 500)
		return
	}
	fmt.Println(limitInt,pageInt)
	jobs := jobService.Jobs(limitInt,pageInt)
	allRecordJob := jobService.JobsCount()
	response.RespondSuccess(c,gin.H{
		"data":jobs,
		"totalPage":allRecordJob/limitInt,
		"currentPage":pageInt,
	},200)
	return
}


func (jobControllers *JobControllers) Job(c *gin.Context){
	jobService := service.JobService{}
	jobId,_:= uuid.FromString(c.Params.ByName("jobId"))
	job,isNotFound:= jobService.FindJobWithID(jobId)
	if(isNotFound){
		response.RespondWithError(c,constant.JOB_NOT_FOUND,500)
		return
	}else{
		response.RespondSuccess(c,gin.H{
			"data":job,
		},200)
		return
	}
	
}
