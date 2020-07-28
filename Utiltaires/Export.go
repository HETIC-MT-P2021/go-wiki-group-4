package Utiltaires
//
messages du package message.go
// MessagePlatform - Comportement pour chaque plateforme
interface de type MessagePlatform {
SendMessage (chaîne de message)
}
// MessagePlatforms - Le registre des plateformes disponibles
var MessagePlatforms = map [chaîne] MessagePlatform {
	"A": PlatformA {Name: "Platform A"},
	"B": PlatformB {Name: "Platform B"},
}