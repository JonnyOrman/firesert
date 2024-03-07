package firesert

func RunMapped[TReceive any, TInsert any](dataMapper DataMapper[TReceive, TInsert]) {
	application := BuildApplicationMapped[TReceive, TInsert](dataMapper)

	application.Run()
}
