#include "net.h"
#include <stdio.h>

void index_handler(http_request *req) {
  printf("hello world\n");
  printf("[%s] %s %s\n", req->method, req->path,req->query);
}

int main(int argc, char *argv[]) {
  http_handle_func("/", index_handler);
  http_listen_and_serve(":8080");
}
