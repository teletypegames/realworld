package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var SecurityACP = DialogManager.Register(&Dialog{
	Name:  "security_acp",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: character.SecurityACP,
			Text: Text{
				{Lang: "en", Text: "UNREGISTERED PRESENCE. STATE YOUR BUSINESS."},
				{Lang: "hu", Text: "NEM REGISZTRÁLT JELENLÉT. KÖZÖLJE A CÉLJÁT."},
			},
		}},
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "Who gave you your orders?"},
					{Lang: "hu", Text: "Kitől kaptad az utasításaidat?"},
				},
				Actions: []Action{
					character.SecurityACP.Say(Text{
						{Lang: "en", Text: "THE OPERATOR. NOBODY ENTERS. THE SIMULATION IS NOT INTERRUPTED. NO TAPE IS CARRIED IN."},
						{Lang: "hu", Text: "AZ OPERÁTOR. SENKI NEM LÉP BE. A SZIMULÁCIÓ NEM SZAKÍTHATÓ MEG. KAZETTA NEM VIHETŐ BE."},
					}),
					character.SecurityACP.Say(Text{
						{Lang: "en", Text: "VOICEPRINT OF THE OPERATOR MATCHES THE SUBJECT IN THE CHAIR. THIS IS NOT A CONTRADICTION."},
						{Lang: "hu", Text: "AZ OPERÁTOR HANGMINTÁJA EGYEZIK A SZÉKBEN ÜLŐ ALANYÉVAL. EZ NEM ELLENTMONDÁS."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "Norman locked himself in. Or somebody with his voice did."},
						{Lang: "hu", Text: "Norman zárta be magát. Vagy valaki az ő hangjával."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Let me talk to Norman."},
					{Lang: "hu", Text: "Hadd beszéljek Normannel."},
				},
				Actions: []Action{
					character.SecurityACP.Say(Text{
						{Lang: "en", Text: "DENIED. THE SUBJECT IS NOT AVAILABLE FOR CONVERSATION. THE SUBJECT IS THE CONVERSATION."},
						{Lang: "hu", Text: "MEGTAGADVA. AZ ALANY NEM ELÉRHETŐ BESZÉLGETÉSRE. AZ ALANY MAGA A BESZÉLGETÉS."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What is the simulation?"},
					{Lang: "hu", Text: "Mi az a szimuláció?"},
				},
				Actions: []Action{
					character.SecurityACP.Say(Text{
						{Lang: "en", Text: "CLASSIFIED. BVK NEURAL INTERFACE, PROTOTYPE TWO. HOST: ONE. UPTIME: CLASSIFIED."},
						{Lang: "hu", Text: "MINŐSÍTETT. BVK NEURÁLIS INTERFÉSZ, KETTES PROTOTÍPUS. GAZDA: EGY. ÜZEMIDŐ: MINŐSÍTETT."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "Host: one. He's not running it. He is it."},
						{Lang: "hu", Text: "Gazda: egy. Nem ő futtatja. Ő maga az."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Raj, this one's for you."},
					{Lang: "hu", Text: "Raj, ez a tiéd."},
				},
				Show: Given{
					Var:     inc.VarSlot2,
					Is:      item.SupportTape,
					NotFlag: inc.FlagACPDistracted,
				},
				Actions: []Action{
					inkwell.RunScript(script.ACPDistracted),
					inkwell.EndDialog(),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Never mind."},
					{Lang: "hu", Text: "Hagyjuk."},
				},
				Actions: []Action{
					character.SecurityACP.Say(Text{
						{Lang: "en", Text: "ACCESS DENIED."},
						{Lang: "hu", Text: "HOZZÁFÉRÉS MEGTAGADVA."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
