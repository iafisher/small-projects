#include <fcntl.h>
#include <semaphore.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/eventfd.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <unistd.h>

const int SEM_INITIAL_VALUE = 4;
const int CHILD_DECREMENT_COUNT = 2;

#define HANDLE_ERR(call) do { if ((call) < 0) { perror(#call); return 1; } } while (0)
#define HANDLE_ERR1(x, call) do { if ((x = (call)) < 0) { perror(#call); return 1; } } while (0)

pid_t parent_pid = 0;
const char* sem_path = "/test-semaphore";

void cleanup() {
  // only run in parent, not in child
  if (getpid() == parent_pid) {
    printf("removing semaphore: %s\n", sem_path);
    sem_unlink(sem_path);
  }
}

int main() {
  parent_pid = getpid();

  sem_unlink(sem_path);
  sem_t* sem = sem_open(sem_path, O_CREAT | O_EXCL, 0600, SEM_INITIAL_VALUE);
  if (sem == SEM_FAILED) {
    perror("sem_open");
    return 1;
  }
  atexit(cleanup);

  int child_to_parent;
  HANDLE_ERR1(child_to_parent, eventfd(0, 0));

  int parent_to_child;
  HANDLE_ERR1(parent_to_child, eventfd(0, 0));

  // any non-zero value works (eventfd blocks if value is 0)
  char buf[8] = { 1 };
  pid_t pid;
  HANDLE_ERR1(pid, fork());
  if (pid == 0) {
    // child
    for (int i = 0; i < CHILD_DECREMENT_COUNT; i++) {
      HANDLE_ERR(sem_wait(sem));
    }

    // 1. Child tells parent that it has waited on semaphore.
    printf("child: called sem_wait %d time(s)\n", CHILD_DECREMENT_COUNT);
    HANDLE_ERR(write(child_to_parent, buf, sizeof buf));

    // 5. Child hears from parent that it has checked semaphore value and can now exit.
    HANDLE_ERR(read(parent_to_child, buf, sizeof buf));

    puts("child: exiting");
    return 0;
  } else {
    // parent

    // 2. Parent hears from child that it has waited on semaphore.
    HANDLE_ERR(read(child_to_parent, buf, sizeof buf));

    int value;
    HANDLE_ERR(sem_getvalue(sem, &value));

    // 3. Parent checks the semaphore's expected value.
    int expected_value = SEM_INITIAL_VALUE - CHILD_DECREMENT_COUNT;
    if (value != expected_value) {
      printf("parent: error: expected semaphore value of %d, got %d\n", expected_value, value);
      return 1;
    }

    // 4. Parent tells child that it has checked semaphore value.
    HANDLE_ERR(write(parent_to_child, buf, sizeof buf));

    // 6. Parent waits for child to exit.
    HANDLE_ERR(waitpid(pid, NULL, 0));

    HANDLE_ERR(sem_getvalue(sem, &value));
    if (value == SEM_INITIAL_VALUE) {
      printf("Semaphore WAS released by child automatically upon exit. (value=%d)\n", value);
    } else if (value == SEM_INITIAL_VALUE - CHILD_DECREMENT_COUNT) {
      printf("Semaphore WAS NOT released by child automatically upon exit. (value=%d)\n", value);
    } else {
      printf("parent: error: got unexpected semaphore value: %d\n", value);
      return 1;
    }
  }

  return 0;
}
