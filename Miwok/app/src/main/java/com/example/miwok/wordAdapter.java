package com.example.miwok;

import android.app.Activity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.ImageView;
import android.widget.TextView;

import java.util.ArrayList;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.core.content.ContextCompat;

public class wordAdapter extends ArrayAdapter<word> {
    private int color_;
    public wordAdapter(Activity context, ArrayList<word> androidFlavors,int color) {
        // Here, we initialize the ArrayAdapter's internal storage for the context and the list.
        // the second argument is used when the ArrayAdapter is populating a single TextView.
        // Because this is a custom adapter for two TextViews and an ImageView, the adapter is not
        // going to use this second argument, so it can be any value. Here, we used 0.
        super(context, 0, androidFlavors);
        color_=color;
    }

    @NonNull
    @Override
    public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
        View listItemView = convertView;
        word currentWord = getItem(position);
        if (listItemView == null) {
            if (currentWord.getImageResourceId() == 0) {
                listItemView = LayoutInflater.from(getContext()).inflate(
                        R.layout.list_item_phrase, parent, false);
                TextView miwokTextView = (TextView) listItemView.findViewById(R.id.id_miwok_phrase);
                miwokTextView.setText(currentWord.getMiwok());
                TextView defaultTextView = (TextView) listItemView.findViewById(R.id.id_translation_phrase);
                defaultTextView.setText(currentWord.getDefault());
                View TextContainer=listItemView.findViewById(R.id.id_translation_phrase);
                View Textcon=listItemView.findViewById(R.id.id_miwok_phrase);
                int col= ContextCompat.getColor(getContext(),color_);
                TextContainer.setBackgroundColor(col);
                Textcon.setBackgroundColor(col);

            } else {
                listItemView = LayoutInflater.from(getContext()).inflate(
                        R.layout.list_item, parent, false);
                TextView miwokTextView = (TextView) listItemView.findViewById(R.id.id_miwok);
                miwokTextView.setText(currentWord.getMiwok());
                TextView defaultTextView = (TextView) listItemView.findViewById(R.id.id_translation);
                defaultTextView.setText(currentWord.getDefault());
                ImageView iconView = (ImageView) listItemView.findViewById(R.id.list_item_icon);
                iconView.setImageResource(currentWord.getImageResourceId());
                View TextContainer=listItemView.findViewById(R.id.id_translation);
                View Textcon=listItemView.findViewById(R.id.id_miwok);
                int col= ContextCompat.getColor(getContext(),color_);
                TextContainer.setBackgroundColor(col);
                Textcon.setBackgroundColor(col);
            }

        }

        return listItemView;
    }
}
