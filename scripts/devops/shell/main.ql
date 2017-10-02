Input = class {

	fn _init() {
		this.scanner = bufio.NewScanner(os.stdin)
	}

	fn Input(prompt) {
		if this.scanner == nil {
			this.scanner = bufio.NewScanner(os.stdin)
		}

		printf("%s> ", prompt)
		scanned = this.scanner.Scan()

		if !scanned {
			return ""
		}

		line = this.scanner.Text()
		
		return line
	}
}
 
Scanner = new Input()

export Scanner