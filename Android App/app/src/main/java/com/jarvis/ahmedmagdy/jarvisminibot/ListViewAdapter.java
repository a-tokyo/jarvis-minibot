package com.jarvis.ahmedmagdy.jarvisminibot;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import com.jarvis.ahmedmagdy.jarvisminibot.Models.ChatBubble;
import com.jarvis.ahmedmagdy.jarvisminibot.Models.NewsChatBubble;
import com.jarvis.ahmedmagdy.jarvisminibot.Models.WeatherChatBubble;

import java.util.ArrayList;

/**
 * Created by ahmedmagdy on 11/29/16.
 */

public class ListViewAdapter  extends BaseAdapter {
    private Context context;
    private LayoutInflater inflater;
    private ArrayList<ChatBubble> dataSource;

    public ListViewAdapter(Context context, ArrayList<ChatBubble> chats) {
        this.context = context;
        this.dataSource = chats;
        this.inflater= (LayoutInflater) this.context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
    }

    @Override
    public int getCount() {
        return dataSource.size();
    }

    @Override
    public Object getItem(int i) {
        return dataSource.get(i);
    }

    @Override
    public long getItemId(int i) {
        return i;
    }

    @Override
    public View getView(int i, View view, ViewGroup viewGroup) {
        View rowView;
        if(dataSource.get(i) instanceof NewsChatBubble){
            rowView = inflater.inflate(R.layout.view_news_bubble, viewGroup, false);
            populateNewsView(rowView, (NewsChatBubble) dataSource.get(i));
        } else{
            rowView = inflater.inflate(R.layout.view_news_bubble, viewGroup, false);
            populateWeatherView(rowView, (WeatherChatBubble) dataSource.get(i));
        }
        return rowView;
    }

    private void populateWeatherView(View view, WeatherChatBubble chatBubble) {
        TextView messagTv = (TextView) view.findViewById(R.id.view_weather_bubble_tv);
        messagTv.setText(chatBubble.getMessage());
    }

    private void populateNewsView(View view, NewsChatBubble chatBubble){
        TextView messagTv = (TextView) view.findViewById(R.id.view_news_bubble_tv);
        messagTv.setText(chatBubble.getMessage());
    }
}
