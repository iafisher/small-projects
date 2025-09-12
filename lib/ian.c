#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "ian.h"

void* ian_malloc(size_t count) {
  void* r = malloc(count);
  if (r == NULL) fatal("out of memory (malloc)");
  return r;
}

void* ian_calloc(size_t count, size_t size) {
  void* r = calloc(count, size);
  if (r == NULL) fatal("out of memory (calloc)");
  return r;
}

void* ian_realloc(void* ptr, size_t size) {
  void* r = realloc(ptr, size);
  if (r == NULL) fatal("out of memory (realloc)");
  return r;
}

void fatal(const char* format, ...) {
  va_list args;
  va_start(args, format);
  fprintf(stderr, "fatal: ");
  vfprintf(stderr, format, args);
  fprintf(stderr, "\n");
  exit(1);
}

char* sfmt(const char* format, ...) {
  va_list args;
  va_start(args, format);
  va_list args2;
  va_copy(args2, args);
  int n = vsnprintf(NULL, 0, format, args);
  va_end(args);
  char* buf = ian_malloc(n + 1);
  vsnprintf(buf, n + 1, format, args2);
  va_end(args2);
  return buf;
}

struct IanStr IanStr_new(const char* s) {
  size_t len = strlen(s);
  char* data = ian_malloc(len + 1);
  memcpy(data, s, len + 1);
  return (struct IanStr){ .len = len, .data = data };
}

void IanStr_concat(struct IanStr* s1, struct IanStr s2) {
  s1->data = ian_realloc(s1->data, s1->len + s2.len + 1);
  memcpy(s1->data + s1->len, s2.data, s2.len + 1);
  s1->len += s2.len;
}

void IanStr_free(struct IanStr s) {
  free(s.data);
}

long long max(long long x, long long y) {
  return x > y ? x : y;
}

const int IAN_STR_BUILDER_INITIAL_SIZE = 256;

struct IanStrBuilder IanStrBuilder_new() {
  char* data = ian_malloc(IAN_STR_BUILDER_INITIAL_SIZE);
  return (struct IanStrBuilder){ .len = 0, .cap = IAN_STR_BUILDER_INITIAL_SIZE, .data = data };
}

void IanStrBuilder_append(struct IanStrBuilder* bldr, const char* s) {
  size_t len = strlen(s);
  size_t new_len = bldr->len + len;
  if (new_len >= bldr->cap) {
    bldr->cap = max(new_len + 1, bldr->cap + IAN_STR_BUILDER_INITIAL_SIZE);
    bldr->data = ian_realloc(bldr->data, bldr->cap);
  }
  memcpy(bldr->data + bldr->len, s, len + 1);
  bldr->len = new_len;
}
