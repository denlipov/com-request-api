package internal

//go:generate mockgen -source=./repo/repo.go -destination=./mocks/repo_mock.go -package=mocks
//go:generate mockgen -source=./app/repo/event.go -destination=./mocks/event_repo_mock.go -package=mocks
//go:generate mockgen -source=./app/sender/event.go -destination=./mocks/event_sender_mock.go -package=mocks
