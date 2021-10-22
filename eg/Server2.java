package shiyan2;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;

public class Server2 {
    public Server2() {
    }

    public static void main(String[] args) throws IOException {
        int i = 0;
        ServerSocket server = new ServerSocket(5602);
        Socket socket = server.accept();
        System.out.println("启动成功");
        OutputStream out = socket.getOutputStream();

        InputStream in;
        for(in = socket.getInputStream(); i < 50; ++i) {
            byte[] bs = new byte[20];
            in.read(bs);
            String str1 = new String(bs);
            out.write(1);
            byte[] ba;
            String str;
            if (str1.startsWith("losePackage")) {
                ba = new byte[20];
                in.read(ba);
                str = new String(ba);
                System.out.println("接收的消息是 " + str);
            } else if (str1.startsWith("loseAck")) {
                ba = new byte[20];
                in.read(ba);
                str = new String(ba);
                System.out.println("接收的消息是 " + str);
                System.out.println("接收的消息是 " + str);
            } else if (str1.startsWith("Wrong")) {
                ba = new byte[20];
                in.read(ba);
                str = new String(ba);
                System.out.println("接收的消息是 " + str + "数据错误");
                System.out.println("接收的消息是 " + str);
            } else if (str1.startsWith("noError")) {
                ba = new byte[20];
                in.read(ba);
                str = new String(ba);
                System.out.println("接收的消息是 " + str);
                String ack = "ACK";
                out.write(ack.getBytes());
            }

            System.out.println("------------");
        }

        in.close();
        out.close();
        socket.close();
        server.close();
    }
}
