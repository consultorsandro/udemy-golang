	// Imprimindo número inteiro a partir de um float
	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	// convetendo int para string
	fmt.Println(("Viagem " + strconv.Itoa(1010)))

	//String para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	///String para booleano mediante estrutura condicional
	b, _ := strconv.ParseBool("True")
	if b {
		fmt.Println("Verdadeiro")

	}

}
