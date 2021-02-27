import java.io.*;
import java.net.*;

class Client {
  public static void main(String[] argv) throws Exception {
    Socket clientSocket = new Socket("localhost", 8081);
    DataOutputStream outToServer = new DataOutputStream(clientSocket.getOutputStream());
    BufferedReader inFromServer =
        new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
    outToServer.writeBytes("5Anybody there?\n");
    String modifiedSentence = inFromServer.readLine();
    System.out.println("Client received from server: " + modifiedSentence);
    clientSocket.close();
  }
}
