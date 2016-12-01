package com.jarvis.ahmedmagdy.jarvisminibot.API;

/**
 * Created by ahmedmagdy on 11/29/16.
 */

public class ChatBody {
    String message;
    public ChatBody(String message){
        this.message = message;
    }
    @Override
    public String toString() {
        return message;
    }
}
