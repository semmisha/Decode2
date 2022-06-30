package env

func EnvText() []byte {

	var EnvText string = "APITOKEN=A12CC37A-9BDD-4466-9610-A822FB935111\n" +
		"APIREQUESTURL=https://support.wagner-auto.ru/api/v3/requests\n"

	return []byte(EnvText)

}
