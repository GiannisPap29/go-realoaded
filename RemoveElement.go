package main

func removeElement(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
