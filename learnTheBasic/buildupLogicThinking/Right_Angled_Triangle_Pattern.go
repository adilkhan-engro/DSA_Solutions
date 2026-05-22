package buildupLogicThinking

func Right_Angled_Triangle_Pattern(n int) {
	for i:=1;i<=n;i++{
		for j:=0;j<i;j++{
			print("*")
		}
		println()
	}

}
