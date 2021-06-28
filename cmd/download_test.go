package cmd

import "testing"

func TestDownloadFile(t *testing.T) {
	err := DownloadFile("https://images.unsplash.com/photo-1521575107034-e0fa0b594529?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1048&q=80")
	if err != nil {
		t.Errorf("DownloadFile() Failed")
	}
}
