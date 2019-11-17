package main

import (
	"github.com/pion/webrtc/v2"
	"fmt"
//	"github.com/pion/webrtc/v2/examples/internal/signal"
)

func main() {
// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	offerPc, _ := webrtc.NewPeerConnection(config)
	answerPc, _ := webrtc.NewPeerConnection(config)

	offerPc.OnICECandidate(func (c *webrtc.ICECandidate){
		if c != nil{
			answerPc.AddICECandidate(c.ToJSON())
		}
	})

	answerPc.OnICECandidate(func (c *webrtc.ICECandidate){
		if c != nil{
			offerPc.AddICECandidate(c.ToJSON())
		}
	})

	offer, _ := offerPc.CreateOffer(nil)
	offerPc.SetLocalDescription(offer)
	answerPc.SetRemoteDescription(offer)

	answer, _ := answerPc.CreateAnswer(nil)
	answerPc.SetLocalDescription(answer)
	offerPc.SetRemoteDescription(answer)	

	offerDataChannel, _ := offerPc.CreateDataChannel("demoChannel", nil)

	offerDataChannel.OnMessage(func(message webrtc.DataChannelMessage){
		fmt.Printf("OnMessage: %s \n", string(message.Data))
	})

	answerPc.OnDataChannel(func(answerDataChannel *webrtc.DataChannel){
		fmt.Printf("OnDataChannel: %s \n", answerDataChannel.Label())

		answerDataChannel.OnOpen(func(){
			answerDataChannel.SendText("Thank you flying Pion!")
		})
	})

}