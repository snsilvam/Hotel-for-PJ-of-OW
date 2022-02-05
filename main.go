//Para el hotel se necesita identificar que tipo de rol tiene el pj, tank o support
package main

import "fmt"

type IRollFactory interface {
	SendRoll()
	GetSenderRoll() ISenderRoll
}

type ISenderRoll interface {
	GetSenderRoll() string
	GetSenderTypeRoll() string
}

type SupportRoll struct {
	//Este es el rol Support
}

func (SupportRoll) SendRoll() {
	fmt.Println("I'm a support")
}
func (SupportRoll) GetSenderRoll() ISenderRoll {
	return SupportRollSender{}
}

type SupportRollSender struct {
}

func (SupportRollSender) GetSenderRoll() string {
	return "Support"
}
func (SupportRollSender) GetSenderTypeRoll() string {
	return "OffHeal"
}

type TankRoll struct {
	//Este es el rol Tank
}

func (TankRoll) SendRoll() {
	fmt.Println("I'm a tank")
}
func (TankRoll) GetSenderRoll() ISenderRoll {
	return TankRollSender{}
}

type TankRollSender struct {
}

func (TankRollSender) GetSenderRoll() string {
	return "Tank"
}
func (TankRollSender) GetSenderTypeRoll() string {
	return "OffTank"
}

func getRollFactory(roll string) (IRollFactory, error) {
	if roll == "Support" {
		return &SupportRoll{}, nil
	}
	if roll == "Tank" {
		return &TankRoll{}, nil
	}
	return nil, fmt.Errorf("No roll")
}
func sendRoll(f IRollFactory) {
	f.SendRoll()
}
func getMethod(f IRollFactory) {
	fmt.Println(f.GetSenderRoll().GetSenderTypeRoll())
}
func Welcome() string {
	return "Hello Welcome to my repository"
}
func main() {
	supportFactory, _ := getRollFactory("Support")
	//tankFactory, _ := getRollFactory("Tank")
	sendRoll(supportFactory)
	//sendRoll(tankFactory)

	getMethod(supportFactory)
	//getMethod(tankFactory)
}
