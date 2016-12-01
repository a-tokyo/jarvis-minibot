package com.jarvis.ahmedmagdy.jarvisminibot;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ListView;
import android.widget.Toast;

import com.jarvis.ahmedmagdy.jarvisminibot.API.ChatBody;
import com.jarvis.ahmedmagdy.jarvisminibot.API.ChatResponse;
import com.jarvis.ahmedmagdy.jarvisminibot.API.ChatbotService;
import com.jarvis.ahmedmagdy.jarvisminibot.API.WelcomeResponse;
import com.jarvis.ahmedmagdy.jarvisminibot.Models.ChatBubble;
import com.jarvis.ahmedmagdy.jarvisminibot.Models.NewsChatBubble;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Objects;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class MainActivity extends AppCompatActivity implements View.OnClickListener, Callback<ChatResponse> {
    private String uuid;
    private Button sendBtn;
    private EditText chatEt;
    private ChatbotService chatbotService;
    private ListView chatBubblesListView;
    private ArrayList<ChatBubble> chatBubbles;
    private ListViewAdapter adapter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        initViews();
        sendWelcomeRequest();
    }

    private void sendWelcomeRequest() {
        chatbotService = ChatbotService.retrofit.create(ChatbotService.class);
        Call<WelcomeResponse> call = chatbotService.welcome();
        call.enqueue(new Callback<WelcomeResponse>() {
            @Override
            public void onResponse(Call<WelcomeResponse> call, Response<WelcomeResponse> response) {
                uuid = response.body().uuid;
                sendBtn.setEnabled(true);
            }

            @Override
            public void onFailure(Call<WelcomeResponse> call, Throwable t) {
                Toast.makeText(getApplicationContext(), t.getMessage(), Toast.LENGTH_LONG).show();

            }
        });

    }

    private void initViews() {
        sendBtn = (Button) findViewById(R.id.chat_send_btn);
        chatEt = (EditText) findViewById(R.id.chat_edit_text);
        chatBubblesListView = (ListView) findViewById(R.id.chat_bubbles_lv);
        sendBtn.setOnClickListener(this);
        chatBubbles = new ArrayList<>();
        adapter = new ListViewAdapter(this, chatBubbles);
        chatBubblesListView.setAdapter(adapter);
    }

    @Override
    public void onClick(View view) {
        String message = chatEt.getText().toString();
        if (uuid != null) {
            if (!message.equals("")) {
                ChatBody chatBody = new ChatBody(message);
                Call<ChatResponse> call = chatbotService.chat(chatBody, uuid);
                call.enqueue(this);
            } else {
                Toast.makeText(getApplicationContext(), "Please write a message first", Toast.LENGTH_LONG).show();
            }
        }
    }


    @Override
    public void onResponse(Call<ChatResponse> call, Response<ChatResponse> response) {
        String message = response.body().toString();
        chatEt.setText("");
        chatBubbles.add(new NewsChatBubble(message));
        adapter.notifyDataSetChanged();
        chatBubblesListView.smoothScrollToPosition(chatBubbles.size());
    }

    @Override
    public void onFailure(Call<ChatResponse> call, Throwable t) {

    }
}
