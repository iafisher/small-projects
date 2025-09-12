#ifndef IAN_H_
#define IAN_H_

#include <stddef.h>
#include <string.h>

void* ian_malloc(size_t);
void* ian_calloc(size_t, size_t);
void* ian_realloc(void*, size_t);

void fatal(const char* format, ...) __attribute__((format(printf, 1, 2)));

char* sfmt(const char* format, ...) __attribute__((format(printf, 1, 2)));

struct IanStr {
  size_t len;
  char* data;
};

struct IanStr IanStr_new(const char*);
void IanStr_concat(struct IanStr*, struct IanStr);
void IanStr_free(struct IanStr);

struct IanStrBuilder {
  size_t len, cap;
  char* data;
};

struct IanStrBuilder IanStrBuilder_new();
void IanStrBuilder_append(struct IanStrBuilder*, const char*);
void IanStrBuilder_fmt(struct IanStrBuilder*, const char* format, ...);

long long max(long long, long long);

#define ian_assert_str_eq(s1, s2) \
  do { \
    const char* ian_s1 = (s1); \
    const char* ian_s2 = (s2); \
    if (strcmp(ian_s1, ian_s2) != 0) { \
      fatal("strings not equal\n  s1: \"%s\"\n  s2: \"%s\"", ian_s1, ian_s2); \
    } \
  } while (0)

#endif
