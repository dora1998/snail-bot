package commands

import (
	"github.com/dora1998/snail-bot/mock_repository"
	"github.com/dora1998/snail-bot/mock_twitter"
	"github.com/dora1998/snail-bot/repository"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestCommandHandler_list(t *testing.T) {
	type args struct {
		username string
		statusId int64
	}

	tests := []struct {
		name     string
		args     args
		injector func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient)
	}{
		{
			"one task",
			args{
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				deadline, _ := time.Parse("2006/1/2 MST", "2019/12/31 JST")

				mr := mock_repository.NewMockRepository(c)
				mr.EXPECT().GetAllTasks().Return([]repository.Task{{
					Id:        "hoge",
					Body:      "task test",
					Deadline:  deadline,
					CreatedAt: time.Time{},
					CreatedBy: "testuser",
				}}).Times(1)

				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().Reply(gomock.Eq("[12/31〆]task test\n"), gomock.Eq(int64(0))).Times(1)
				return mr, mt
			},
		},
		{
			"no task",
			args{
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				mr := mock_repository.NewMockRepository(c)
				mr.EXPECT().GetAllTasks().Return([]repository.Task{}).Times(1)

				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().Reply(gomock.Eq("現在出ている課題はありません🎉"), gomock.Eq(int64(0))).Times(1)
				return mr, mt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo, twitterClient := tt.injector(ctrl)
			h := &CommandHandler{
				repository:    repo,
				twitterClient: twitterClient,
			}
			h.list(tt.args.statusId)
		})
	}
}
