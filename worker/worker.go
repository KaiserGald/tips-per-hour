// Package worker
// 6 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg
//
// Package worker contains the worker struct, and the methods for working with the worker struct.

package worker

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"

	"golang.org/x/crypto/bcrypt"
)

// Worker is a struct contianing a user name and password
type Worker struct {
	name string
	pass string
}

// HashPassword hashes and sets the worker password
func (w *Worker) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	w.pass = string(bytes)
	return err
}

// CheckPasswordHash checks the password to see if it is valid
func (w *Worker) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(w.pass), []byte(password))
	return err == nil
}

// SetWorkerName sets the workername to the given name
func (w *Worker) SetWorkerName(wn string) {
	w.name = wn
}

// GetWorkerName gets the worker's name
func (w *Worker) GetWorkerName() string {
	return w.name
}

// Save saves worker data to disk
func (w *Worker) Save() error {

	// Get home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// If directory doesn't exist, create new directory
	if _, err = os.Stat(usr.HomeDir + "/.tph/workers/" + w.name); err != nil {
		if os.IsNotExist(err) {
			exec.Command("mkdir", usr.HomeDir+"/.tph/workers/"+w.name).Run()
		}
	}

	// Write data to file
	d := []byte(w.name + "\n" + w.pass + "\n")
	err = ioutil.WriteFile(usr.HomeDir+"/.tph/workers/"+w.name+"/"+w.name, d, 0644)
	if err != nil {
		return err
	}

	log.Printf("Worker %v saved in ~/.tph/workers/%v.\n", w.name, w.name)

	return nil
}
