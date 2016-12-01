package com.jarvis.ahmedmagdy.jarvisminibot.Models;

/**
 * Created by ahmedmagdy on 11/29/16.
 */

public class WeatherChatBubble extends ChatBubble {
    private String message;

    public WeatherChatBubble(String message) {
        this.message = message;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
