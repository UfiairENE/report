package errorMsg

import "github.com/rs/zerolog/log"

const ErrorInvalidSecret = "Invalid secret key provided.\nKindly provide a correct secret to continue"
const ErrorNotFunded = "Account does not exist.\nKindly fund account to continue"
const ErrorSessionNotSaved = "Failed to save session"
const ErrorLoginFirst = "Please login first"
const ErrorLogutFirst = "Please logout first"
const ErrorLoginContinue = "Please login again to continue"
const ErrorAccountCheck = "Account different with the sign account"
const ErrorTnxNotBuild = "Transcation failed to build. \nTry again later"
const ErrorTnxNotSign = "Transcation failed to sign. \n Make your account is correct"
const ErrorTnxB64Failed = "Transcation failed to build to base64"

// https://betterstack.com/community/guides/logging/zerolog/

func LogError(message interface{}) {

	//log.Logger = log.With().Caller().Logger()

	log.Error().Msgf("%v", message)

}

func LogInfo(message interface{}) {

	//log.Logger = log.With().Caller().Logger()

	//log.Info().Str("ip","898989").Msg("sign up")

	log.Info().Msgf("%v ", message)

}

func LogDebug(message interface{}) {

	//log.Logger = log.With().Caller().Logger()

	log.Debug().Msgf("%v ", message)

}
