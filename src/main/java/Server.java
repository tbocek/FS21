import java.net.*;
import java.nio.charset.StandardCharsets;

class Server {
  public static void main(String args[]) throws Exception {
    DatagramSocket serverSocket = new DatagramSocket(8081);
    byte[] receiveData = new byte[1024];
    while (true) {
      DatagramPacket receivePacket = new DatagramPacket(receiveData, receiveData.length);
      serverSocket.receive(receivePacket);
      String s = new String(receivePacket.getData()).trim();
      System.out.println("Message Received: " + s);
      InetAddress address = receivePacket.getAddress();
      int port = receivePacket.getPort();
      byte[] reply = ("Message Received: " + s.toUpperCase()).getBytes(StandardCharsets.UTF_8);
      DatagramPacket sendPacket = new DatagramPacket(reply, reply.length, address, port);
      serverSocket.send(sendPacket);
      receiveData = new byte[1024];
    }
  }
}
