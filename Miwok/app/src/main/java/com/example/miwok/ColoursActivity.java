package com.example.miwok;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.ListView;

import java.util.ArrayList;

public class ColoursActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_colours);
        ArrayList<word> colours=new ArrayList<word>();
        colours.add(new word("weṭeṭṭi","red",R.drawable.color_red));
        colours.add(new word("chokokki","green",R.drawable.color_green));
        colours.add(new word("ṭakaakki","brown",R.drawable.color_brown));
        colours.add(new word("ṭopoppi","gray",R.drawable.color_gray));
        colours.add(new word("kululli","black",R.drawable.color_black));
        colours.add(new word("kelelli","white",R.drawable.color_white));
        colours.add(new word("ṭopiisә","dusty yellow",R.drawable.color_dusty_yellow));
        colours.add(new word("ṭopiisә","mustard yellow",R.drawable.color_mustard_yellow));


        wordAdapter itemsAdapter = new wordAdapter(this,colours,R.color.category_colors);
        ListView listView = (ListView) findViewById(R.id.list_colour);

        listView.setAdapter(itemsAdapter);

    }
}