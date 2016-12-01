package com.jarvis.ahmedmagdy.jarvisminibot.API;

import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.Header;
import retrofit2.http.POST;
import retrofit2.http.Path;

/**
 * Created by ahmedmagdy on 11/29/16.
 */

public interface ChatbotService {
    @GET("/welcome")
    Call<WelcomeResponse> welcome();

    @POST("/chat")
    Call<ChatResponse> chat(@Body ChatBody body, @Header("Authorization") String uuid);

    public static final Retrofit retrofit = new Retrofit.Builder()
            .baseUrl("https://jarvis-minibot.herokuapp.com/")
            .addConverterFactory(GsonConverterFactory.create())
            .build();
}
