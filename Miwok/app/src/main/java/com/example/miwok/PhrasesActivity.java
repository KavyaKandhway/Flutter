package com.example.miwok;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.ListView;

import java.util.ArrayList;

public class PhrasesActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_phrases);
        ArrayList<word> phrases=new ArrayList<word>();
        phrases.add(new word("minto wuksus","Where are you going?",0));
        phrases.add(new word("tinnә oyaase'nә","What is your name?",0));
        phrases.add(new word("oyaaset..","My name is..",0));
        phrases.add(new word("michәksәs?","How are you feeling?",0));
        phrases.add(new word("kuchi achit","I’m feeling good.",0));
        phrases.add(new word("әәnәs'aa?","Are you coming?",0));
        phrases.add(new word("hәә’ әәnәm","Yes, I’m coming",0));
        phrases.add(new word("әәnәm","I’m coming",0));
        phrases.add(new word("yoowutis","Let’s go",0));
        phrases.add(new word("әnni'nem","Come here.",0));

        wordAdapter itemsAdapter = new wordAdapter(this,phrases,R.color.category_phrases);
        ListView listView = (ListView) findViewById(R.id.list_phrase);

        listView.setAdapter(itemsAdapter);

    }
}