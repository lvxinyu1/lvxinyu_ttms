package model

type Seat struct {
	Flag int
	Row int
	Line int
}
type Sit struct {
	PlanId string
	Perf   *Plan
	SitOne []Seat
	SitTwo []Seat
	SitThree []Seat
	SitFour []Seat
	SitFive []Seat
	SitSix []Seat
	SitSeven []Seat

}
