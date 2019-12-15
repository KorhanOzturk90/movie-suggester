package main

import (
	"fmt"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkExtractMovieIdFromTitleLink(*testing.T) {
	extractedMovieId := extractMovieIdFromTitleLink("/title/tt0120586/?ref_=tt_rec_tti")
	fmt.Println("extractor result: ", extractedMovieId)
}

func TestAlexaHandler(t *testing.T) {

	intentSlots:= make(map[string]alexa.IntentSlot)
	intentSlots["movie"] = alexa.IntentSlot{
		Name:"movie",
		Value:"Friends with benefits",
	}
	intent := alexa.Intent{
		Name: Recommended_movie_intent,
		Slots: intentSlots,
	}

	request := alexa.Request{
		Intent: intent,
		Type: "IntentRequest",
	}

	att := alexa.Session{}.Attributes

	session := &alexa.Session{
		SessionID:  "testId",
		Attributes: att,
	}

	requestEnv := alexa.RequestEnvelope{
		Request: &request,
		Session: session,

	}

	sendAlexaCommand(&requestEnv)
}

func TestAlexaTopStreamingMoviesHandler(t *testing.T) {

	intentSlots:= make(map[string]alexa.IntentSlot)

	intent := alexa.Intent{
		Name: Recommended_streaming_intent,
		Slots: intentSlots,
	}
	request := alexa.Request{
		Intent: intent,
		Type: "IntentRequest",
	}

	att := alexa.Session{}.Attributes

	session := &alexa.Session{
		SessionID:  "testId",
		Attributes: att,
	}

	requestEnv := alexa.RequestEnvelope{
		Request: &request,
		Session: session,
	}

	sendAlexaCommand(&requestEnv)
}

func sendAlexaCommand(requestEnv *alexa.RequestEnvelope) {
	response, err := Handle(nil, requestEnv)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("response: %v", response.(*alexa.ResponseEnvelope).Response.OutputSpeech.Text )
	}
}


func TestAlexaHandlerNoMovieSpecified(t *testing.T) {
	intentMap := make(map[string]alexa.IntentSlot)
	intent := alexa.Intent{"movieparserIntent", "", intentMap}
	alexaRequest := alexa.Request{"", "", "", "", "", intent, "movie suggester"}
	outputSpeech := alexa.OutputSpeech{"", "", ""}
	card := alexa.Card{"", "", "", "", nil}
	alexaResponse := alexa.Response{&outputSpeech, &card, nil, nil, true}
	processAlexaIntent(&alexaRequest, &alexaResponse)
	assert.Equal(t, "Please make sure you specify the movie name based on which recommendations will be made", alexaResponse.OutputSpeech.Text, "Error message should be returned when there is no movie name!")
}

func TestExtractMovieTitleFromLinkRegex(t *testing.T) {
	extractedTitle := extractMovieIdFromTitleLink("/title/tt0071562/?ref_=tt_sims_tti")
	assert.Equal(t, "tt0071562", extractedTitle, "Unsuccessful movie title retrieval from link")
}

func TestRedisClient(t *testing.T) {
	redisClient()
}
