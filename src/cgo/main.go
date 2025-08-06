package main

/*
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <stdlib.h>
#include <string.h>

#define SHM_SIZE 1024  // Size of the shared memory

key_t get_key() {
    return ftok("/tmp", 65);  // Generate a unique key
}

int create_shm(key_t key) {
    return shmget(key, SHM_SIZE, 0666 | IPC_CREAT);
}

char* attach_shm(int shmid) {
    return (char*) shmat(shmid, NULL, 0);
}

void detach_shm(char *shmaddr) {
    shmdt(shmaddr);
}

void remove_shm(int shmid) {
    shmctl(shmid, IPC_RMID, NULL);
}
*/
import (
	"C"
)
import (
	"fmt"
)

func main() {
	// Generate a key
	key := C.get_key()
	if key == -1 {
		fmt.Println("Failed to generate key")
		return
	}

	// Create shared memory
	shmid := C.create_shm(key)
	if shmid == -1 {
		fmt.Println("Failed to create shared memory")
		return
	}

	fmt.Printf("Shared memory created with ID: %d\n", shmid)

	// Attach shared memory
	shmaddr := C.attach_shm(shmid)
	if shmaddr == nil {
		fmt.Println("Failed to attach shared memory")
		return
	}
	fmt.Println("Shared memory attached")

	// Write data to shared memory
	message := "Hello, Shared Memory!"
	C.strcpy((*C.char)(shmaddr), C.CString(message))
	fmt.Println("Data written to shared memory")

	// Read data from shared memory
	goData := C.GoString((*C.char)(shmaddr))
	fmt.Printf("Read from shared memory: %s\n", goData)

	// Detach shared memory
	C.detach_shm(shmaddr)
	fmt.Println("Shared memory detached")

	// Remove shared memory (optional)
	C.remove_shm(shmid)
	fmt.Println("Shared memory removed")
}
