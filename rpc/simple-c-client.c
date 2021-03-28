#include <stdio.h>
#include <sys/socket.h>
#include <string.h>
#include <arpa/inet.h>

int main() {
    int sock;
    char buffer[16];
    struct sockaddr_in server;

    //Create socket
    sock = socket(AF_INET, SOCK_STREAM, 0);

    server.sin_family = AF_INET;
    server.sin_port = 7000;
    server.sin_addr.s_addr = inet_addr("127.0.0.1");

    //Connect
    connect(sock, (struct sockaddr *) &server, sizeof(server));

    //Send
    buffer[0] = 5;
    strcpy(buffer + 1, "Anybody there?");
    send(sock, buffer, 16, 0);

    //Cleanup and close... todo
    return 0;
}