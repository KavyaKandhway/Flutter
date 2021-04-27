package com.example.miwok;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.ListView;

import java.util.ArrayList;

public class FamilyMembersActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_family_members);
        ArrayList<word> family=new ArrayList<word>();
        family.add(new word("әpә","father",R.drawable.family_father));
        family.add(new word("әṭa","mother",R.drawable.family_mother));
        family.add(new word("angsi","son",R.drawable.family_son));
        family.add(new word("tune","daughter",R.drawable.family_daughter));
        family.add(new word("taachi","older brother",R.drawable.family_older_brother));
        family.add(new word("chalitti","younger brother",R.drawable.family_younger_brother));
        family.add(new word("teṭe","older sister",R.drawable.family_older_sister));
        family.add(new word("kolliti","younger sister",R.drawable.family_younger_sister));
        family.add(new word("ama","grandmother",R.drawable.family_grandmother));
        family.add(new word("paapa","grandfather",R.drawable.family_grandfather));

        wordAdapter itemsAdapter = new wordAdapter(this,family,R.color.category_family);
        ListView listView = (ListView) findViewById(R.id.list_family);

        listView.setAdapter(itemsAdapter);
    }
}