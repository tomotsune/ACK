package shiyan2;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;
import java.net.UnknownHostException;
import java.util.Random;

public class Client2 {
    int message;
    Random random;
    String info;
    int rdNum;

    public Client2(int msg) {
        this.message = msg;
        this.MsgState();
    }

    public void MsgState() {
        this.random = new Random();
        this.rdNum = Math.abs(this.random.nextInt(100));
        if (this.rdNum < 20) {
            this.info = "losePackage";
        } else if (this.rdNum < 40 && this.rdNum >= 20) {
            this.info = "loseAck";
        } else if (this.rdNum < 60 && this.rdNum >= 40) {
            this.info = "Wrong";
        } else if (this.rdNum >= 60) {
            this.info = "noError";
        }

    }
// 客户端发送的数据，服务端接收到了就在原来的基础上+1，服务端在返回给客户端
    public int getMsg() {
        return this.message;
    }

    public void setMsg(int msg) {
        this.message = msg;
    }

    public String getInfo() {
        return this.info;
    }

    public static void main(String[] args) throws UnknownHostException, IOException {
        Socket socket = new Socket("localhost", 5602);
        OutputStream out = socket.getOutputStream();
        InputStream in = socket.getInputStream();
        Random random = new Random();

        for(int i = 0; i < 50; ++i) {
            int number = random.nextInt(255) % 240 + 16;
            Client2 msg = new Client2(number);
            String str = Integer.toBinaryString(msg.getMsg());
            System.out.println("发送的信息：" + str);
            System.out.println("发送状态： " + msg.getInfo());
            out.write(msg.getInfo().getBytes());
            byte[] ba = new byte[100];
            in.read(ba);
            if (msg.getInfo().startsWith("losePackage")) {
                System.out.println("到时重发");
                out.write(str.getBytes());
            } else if (msg.getInfo().startsWith("loseAck")) {
                System.out.println("到时重发");
                out.write(str.getBytes());
            } else if (msg.getInfo().startsWith("Wrong")) {
                System.out.println("到时重发");
                out.write(str.getBytes());
            } else if (msg.getInfo().startsWith("noError")) {
                out.write(str.getBytes());
                byte[] bs = new byte[100];
                in.read(bs);
                System.out.println("ACK");
            }

            System.out.println("------------");
        }

        in.close();
        out.close();
        socket.close();
    }
}
