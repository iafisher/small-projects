#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "ian.h"


int LOGGING = 0;


struct AvlNode {
  const char* key;
  const char* value;
  int balance_factor;
  struct AvlNode *left, *right;
};

void print_avl_tree(const struct AvlNode *root);

struct AvlNode* AvlNode_create(const char* key, const char* value) {
  struct AvlNode* r = calloc(1, sizeof *r);
  assert(r != NULL);
  r->key = key;
  r->value = value;
  return r;
}

struct AvlNode_insertion {
  struct AvlNode* node;
  unsigned int height_change;
};

unsigned int height(struct AvlNode* r) {
  if (r == NULL) {
    return 0;
  } else {
    unsigned int left_height = height(r->left);
    unsigned int right_height = height(r->right);
    return (left_height > right_height ? left_height : right_height) + 1;
  }
}

void check_balance_factor(struct AvlNode* r) {
  if (r == NULL) return;

  check_balance_factor(r->left);
  check_balance_factor(r->right);
  unsigned int left_height = height(r->left);
  unsigned int right_height = height(r->right);
  if (left_height - right_height != r->balance_factor) {
    printf("ERROR: balance factor should be %d (%d - %d), actually is %d\n",
        left_height - right_height,
        left_height,
        right_height,
        r->balance_factor);
    print_avl_tree(r);
    puts("");
  }
}

void check_binary_tree(struct AvlNode* r) {
  if (r == NULL) return;

  if (r->left) {
    int cmp = strcmp(r->left->key, r->key);
    if (cmp >= 0) {
      printf("ERROR: binary tree not sorted on left (strcmp(\"%s\", \"%s\") == %d)\n", r->left->key, r->key, cmp);
      print_avl_tree(r);
      puts("");
    }
  }

  if (r->right) {
    int cmp = strcmp(r->key, r->right->key);
    if (cmp >= 0) {
      printf("ERROR: binary tree not sorted on right (strcmp(\"%s\", \"%s\") == %d)\n", r->key, r->right->key, cmp);
      print_avl_tree(r);
      puts("");
    }
  }

  /* check_binary_tree(r->left); */
  /* check_binary_tree(r->right); */
}

struct AvlNode* AvlNode_rotate_right(struct AvlNode* r) {
  //       r
  //      / \
  //     c   T3
  //    / \
  //   T1  T2
  //   |
  //   n
  //
  //   becomes
  //
  //        c
  //      /   \
  //     T1    r
  //     |    / \
  //     n   T2  T3
  /* puts("BEFORE ROTATE RIGHT"); */
  /* print_avl_tree(r); */
  struct AvlNode* c = r->left;
  struct AvlNode* t2 = c->right;

  c->right = r;
  r->left = t2;
  /* puts("AFTER ROTATE RIGHT"); */
  /* print_avl_tree(c); */
  return c;
}

struct AvlNode* AvlNode_rotate_left(struct AvlNode* root) {
  /* puts("BEFORE ROTATE LEFT"); */
  /* print_avl_tree(root); */
  struct AvlNode* right = root->right;
  struct AvlNode* right_left = right->left;

  right->left = root;
  root->right = right_left;
  /* puts("AFTER ROTATE LEFT"); */
  /* print_avl_tree(right); */
  return right;
}

struct AvlNode* AvlNode_rotate_left_right(struct AvlNode* root) {
  int bf = root->left->right->balance_factor;
  /* puts("BEFORE ROTATE LEFT-RIGHT"); */
  /* print_avl_tree(root); */
  root->left = AvlNode_rotate_left(root->left);
  struct AvlNode* ret = AvlNode_rotate_right(root);
  if (bf == 0) {
    ret->left->balance_factor = 0;
    ret->right->balance_factor = 0;
  } else if (bf == 1) {
    ret->right->balance_factor = 0;
    ret->left->balance_factor = -1;
  } else {
    ret->right->balance_factor = 1;
    ret->left->balance_factor = 0;
  }
  /* puts("AFTER ROTATE LEFT-RIGHT"); */
  /* print_avl_tree(ret); */
  return ret;
}

struct AvlNode* AvlNode_rotate_right_left(struct AvlNode* root) {
  int bf = root->right->left->balance_factor;
  /* puts("BEFORE ROTATE RIGHT-LEFT"); */
  /* print_avl_tree(root); */
  root->right = AvlNode_rotate_right(root->right);
  struct AvlNode* ret = AvlNode_rotate_left(root);
  ret->balance_factor = 0;
  if (bf == 0) {
    ret->left->balance_factor = 0;
    ret->right->balance_factor = 0;
  } else if (bf == 1) {
    ret->left->balance_factor = 0;
    ret->right->balance_factor = -1;
  } else {
    ret->left->balance_factor = 1;
    ret->right->balance_factor = 0;
  }
  /* puts("AFTER ROTATE RIGHT-LEFT"); */
  /* print_avl_tree(ret); */
  return ret;
}

struct AvlNode_insertion AvlNode_insert_helper(struct AvlNode* root, const char* key, const char* value) {
  if (root == NULL) {
    if (LOGGING) puts("inserting new leaf");
    return (struct AvlNode_insertion){ .node = AvlNode_create(key, value), .height_change = 1 };
  }

  int cmp = strcmp(key, root->key);
  unsigned int height_change;
  if (cmp == 0) {
    root->value = value;
    return (struct AvlNode_insertion){ .node = root, .height_change = 0 };
  } else if (cmp < 0) {
    struct AvlNode_insertion insertion = AvlNode_insert_helper(root->left, key, value);
    root->left = insertion.node;
    root->balance_factor += insertion.height_change;
    height_change = root->balance_factor > 0 ? 1 : 0;
    if (LOGGING) {
      printf("after inserting left (height_change=%d)\n", height_change);
      print_avl_tree(root);
    }
  } else {
    struct AvlNode_insertion insertion = AvlNode_insert_helper(root->right, key, value);
    root->right = insertion.node;
    root->balance_factor -= insertion.height_change;
    height_change = root->balance_factor < 0 ? 1 : 0;
    if (LOGGING) {
      printf("after inserting right (height_change=%d)\n", height_change);
      print_avl_tree(root);
    }
  }

  struct AvlNode* new_root = root;
  if (root->balance_factor == 2) {
    if (root->left->balance_factor == 1) {
      new_root = AvlNode_rotate_right(root);
      new_root->balance_factor = 0;
      new_root->right->balance_factor = 0;
      if (LOGGING) {
        puts("rotate_right");
        print_avl_tree(new_root);
      }
    } else {
      new_root = AvlNode_rotate_left_right(root);
      if (LOGGING) {
        puts("rotate_left_right");
        print_avl_tree(new_root);
      }
    }
    /* check_balance_factor(new_root); */
    height_change = 0;
  } else if (root->balance_factor == -2) {
    if (root->right->balance_factor == -1) {
      new_root = AvlNode_rotate_left(root);
      new_root->balance_factor = 0;
      new_root->left->balance_factor = 0;
      if (LOGGING) {
        puts("rotate_left");
        print_avl_tree(new_root);
      }
    } else {
      new_root = AvlNode_rotate_right_left(root);
      if (LOGGING) {
        puts("rotate_right_left");
        print_avl_tree(new_root);
      }
    }
    /* check_balance_factor(new_root); */
    height_change = 0;
  }

  return (struct AvlNode_insertion){ .node = new_root, .height_change = height_change };
}

struct AvlNode* AvlNode_insert(struct AvlNode* root, const char* key, const char* value) {
  if (LOGGING) {
    puts("BEFORE");
    print_avl_tree(root);
  }
  struct AvlNode_insertion insertion = AvlNode_insert_helper(root, key, value);
  check_balance_factor(insertion.node);
  check_binary_tree(insertion.node);
  return insertion.node;
}

void print_avl_subtree(const struct AvlNode *node, const char *prefix, int is_last) {
  if (!node) return;

  // Draw connector for this node
  printf("%s%s── %s  %d\n",
         prefix,
         is_last ? "└" : "├",
         node->key ? node->key : "(null)",
         node->balance_factor);

  // Build prefix for this node's children
  char next_prefix[1024];
  snprintf(next_prefix, sizeof next_prefix, "%s%s",
           prefix, is_last ? "    " : "│   ");

  if (node->left == NULL && node->right == NULL) {
    return;
  }

  // Determine which child is "last" to choose ├ vs └ for left
  if (node->right) {
    int right_is_last = (node->left == NULL);
    print_avl_subtree(node->right, next_prefix, right_is_last);
  } else {
    printf("%s%s──\n", next_prefix, node->left ? "├" : "└");
  }
  if (node->left) {
    print_avl_subtree(node->left, next_prefix, 1);
  } else {
    printf("%s└──\n", next_prefix);
  }
}

void print_avl_tree(const struct AvlNode *root) {
  if (!LOGGING) return;

  if (!root) {
    puts("(empty)");
    return;
  }
  // Print root without a leading connector
  printf("%s  %d\n", root->key ? root->key : "(null)", root->balance_factor);

  if (root->left == NULL && root->right == NULL) {
    return;
  }

  // Print children
  if (root->right) {
    int right_is_last = (root->left == NULL);
    print_avl_subtree(root->right, "", right_is_last);
  } else {
    printf("└──\n");
  }
  if (root->left) {
    print_avl_subtree(root->left, "", 1);
  } else {
    printf("└──\n");
  }
}

void avl_tree_to_sexp_builder(struct AvlNode* root, struct IanStrBuilder* bldr) {
  if (root == NULL) {
    IanStrBuilder_append(bldr, "()");
    return;
  }

  // calls to `sfmt` leak memory...

  if (root->left == NULL && root->right == NULL) {
    IanStrBuilder_append(bldr, sfmt("%s:%d", root->key, root->balance_factor));
    return;
  }

  IanStrBuilder_append(bldr, sfmt("(%s:%d ", root->key, root->balance_factor));
  avl_tree_to_sexp_builder(root->left, bldr);
  IanStrBuilder_append(bldr, " ");
  avl_tree_to_sexp_builder(root->right, bldr);
  IanStrBuilder_append(bldr, ")");
}

char* avl_tree_to_sexp(struct AvlNode* root) {
  struct IanStrBuilder bldr = IanStrBuilder_new();
  avl_tree_to_sexp_builder(root, &bldr);
  return bldr.data;
}

/* Size 5: mostly left-leaning with one right child */
struct AvlNode* build_tree_5(void) {
  //        M
  //      /   \
  //     C     T
  //    /
  //   A
  //    \
  //     B
  struct AvlNode* M = AvlNode_create("M", "13");
  struct AvlNode* C = AvlNode_create("C", "3");
  struct AvlNode* T = AvlNode_create("T", "20");
  struct AvlNode* A = AvlNode_create("A", "1");
  struct AvlNode* B = AvlNode_create("B", "2");

  M->left = C;  M->right = T;
  C->left = A;
  A->right = B;
  return M;
}

/* Size 6: mixed shape with both subtrees having depth */
struct AvlNode* build_tree_6(void) {
  //        H
  //      /   \
  //     D     P
  //    / \   /
  //   B   F N
  //        \
  //         G
  struct AvlNode* H = AvlNode_create("H", "8");
  struct AvlNode* D = AvlNode_create("D", "4");
  struct AvlNode* P = AvlNode_create("P", "16");
  struct AvlNode* B = AvlNode_create("B", "2");
  struct AvlNode* F = AvlNode_create("F", "6");
  struct AvlNode* N = AvlNode_create("N", "14");
  struct AvlNode* G = AvlNode_create("G", "7");

  H->left = D; H->right = P;
  D->left = B; D->right = F;
  P->left = N;
  F->right = G;
  return H;
}

/* Size 7: perfectly balanced BST-shaped layout for clarity */
struct AvlNode* build_tree_7(void) {
  //         D
  //       /   \
  //      B     F
  //     / \   / \
  //    A   C E   G
  struct AvlNode* root = AvlNode_create("D", "4");
  root = AvlNode_insert(root, "B", "2");
  root = AvlNode_insert(root, "F", "6");
  root = AvlNode_insert(root, "A", "1");
  root = AvlNode_insert(root, "C", "3");
  root = AvlNode_insert(root, "E", "5");
  root = AvlNode_insert(root, "G", "7");
  return root;
}

void test_insert() {
  struct AvlNode* root = build_tree_7();
  ian_assert_str_eq(avl_tree_to_sexp(root), "(D:0 (B:0 A:0 C:0) (F:0 E:0 G:0))");
}

int main() {
  if (1) {
    test_insert();
    puts("tests passed");
    return 0;
  }
  /* char a_to_z[27]; */
  /* for (size_t i = 0; i < 26; i++) { */
  /*   a_to_z[i] = 'a' + i; */
  /* } */
  /* a_to_z[26] = '\0'; */

  /* srand((unsigned)time(NULL)); */
  /* for (size_t i = 25; i > 0; i--) { */
  /*   size_t j = rand() % (i + 1); */
  /*   char c = a_to_z[i]; */
  /*   a_to_z[i] = a_to_z[j]; */
  /*   a_to_z[j] = c; */
  /* } */

  char* a_to_z = "jmpgbcrziqksatdhuyfenxovlw";
  printf("%s\n\n", a_to_z);

  /* size_t len = 26; */
  size_t len = 8;

  struct AvlNode* root = NULL;
  for (size_t i = 0; i < len; i++) {
    char* s = malloc(2);
    s[0] = a_to_z[i];
    s[1] = '\0';
    if (s[0] == 'z') {
      LOGGING = 1;
    }
    root = AvlNode_insert(root, s, "");
  }

  puts("");
  print_avl_tree(root);

  /* struct AvlNode* root = AvlNode_create("30", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "20", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "10", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "40", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "50", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "25", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "35", ""); */
  /* /1* print_avl_tree(root); *1/ */
  /* /1* puts(""); *1/ */
  /* root = AvlNode_insert(root, "36", ""); */

  /* puts(""); */
  /* puts("FINAL"); */
  /* print_avl_tree(root); */
  /* struct AvlNode* t5 = build_tree_5(); */
  /* struct AvlNode* t6 = build_tree_6(); */
  /* struct AvlNode* t7 = build_tree_7(); */

  /* puts("Tree (size 5):"); */
  /* print_avl_tree(t5); */
  /* puts("\nTree (size 6):"); */
  /* print_avl_tree(t6); */
  /* puts("\nTree (size 7):"); */
  /* print_avl_tree(t7); */
  return 0;
}
