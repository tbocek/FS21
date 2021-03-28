#include <stdio.h>
#include <sys/socket.h>
#include <string.h>
#include <arpa/inet.h>

int main() {
    int socket_desc, client_sock, addr_size, read_size;
    char buffer[16];
    struct sockaddr_in server, client;

    //Create socket
    socket_desc = socket(AF_INET, SOCK_STREAM, 0);

    server.sin_family = AF_INET;
    server.sin_port = 7000;
    server.sin_addr.s_addr = inet_addr("127.0.0.1");

    //Bind
    bind(socket_desc, (struct sockaddr *) &server, sizeof(server));
    listen(socket_desc, 1);

    //Accept
    addr_size = sizeof(struct sockaddr_in);
    client_sock = accept(socket_desc, (struct sockaddr *) &client, &addr_size);

    //Receive
    read_size = recv(client_sock, buffer, 16, 0);
    printf("%d, %s", buffer[0], buffer + 1);
    //Cleanup and close... todo
    return 0;
}