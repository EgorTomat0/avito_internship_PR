package repository

import (
	"log"

	"gorm.io/gorm"
)

func Test(db *gorm.DB) {
	repo := NewPullRequestRepository(db)
	//err := repo.Create(models.PullRequestDomain{
	//	PullRequestID:   "f1be0b61-2d74-4c3c-aa38-ecf77cdb2642",
	//	PullRequestName: "testPR",
	//	AuthorID:        "ed195c47-1576-4d51-968a-42ba435bb488",
	//})
	//pr, err := repo.Merge("f1be0b61-2d74-4c3c-aa38-ecf77cdb2642")
	//pr, err := repo.ReassignReviewer("f1be0b61-2d74-4c3c-aa38-ecf77cdb2642", "057b6790-1355-4ce2-910e-8d7ef3692413")
	pr, err := repo.GetPullRequestsByReviewerUserID("d81c9d6f-2aaa-44f4-bdbd-17212dcd3141")
	if err != nil {
		panic(err)
	}
	log.Println(pr)
}
