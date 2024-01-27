package job

import (
	"fmt"
	"log"
	"time"

	"github.com/rsdel2007/proj/pkg/youtube"
	"gorm.io/gorm"
)

type JobQueue struct {
	jobs          chan Job
	dbInitializer *DatabaseInitializer
}

type Job interface {
	Execute()
}

func NewJobQueue(db *gorm.DB) *JobQueue {
	dbInitializer := NewDatabaseInitializer(db)

	return &JobQueue{
		jobs:          make(chan Job),
		dbInitializer: dbInitializer,
	}
}

func (jq *JobQueue) Enqueue(job Job) {
	jq.jobs <- job
}

func (jq *JobQueue) StartFetchingJob() {
	jq.dbInitializer.Initialize()

	// Start a goroutine to enqueue jobs periodically
	go func() {
		for {
			job := NewFetchJob(jq.dbInitializer.DB)
			jq.Enqueue(job)
			fmt.Println("[job.Fetcher]Enqueue")
			time.Sleep(60 * 1 * time.Second)
		}
	}()

	// Start another goroutine to process jobs
	go func() {
		for {
			select {
			case job := <-jq.jobs:
				fmt.Printf("[job.Fetcher]Dequeued a job\n")
				job.Execute()
			}
		}
	}()
}

type FetchJob struct {
	db *gorm.DB
}

func NewFetchJob(db *gorm.DB) *FetchJob {
	return &FetchJob{
		db: db,
	}
}

// Execute executes the fetch job
func (j *FetchJob) Execute() {
	log.Println("[job.fetcher] Executing job")
	youtube.FetchAndStoreVideos(j.db)
	log.Println("YouTube videos fetched and stored successfully.")
}
