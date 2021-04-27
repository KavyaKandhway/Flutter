package com.example.miwok;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
    public void openNumberActivity(View view)
    {
        Intent i = new Intent(this, NumbersActivity.class);
        startActivity(i);

    }
    public void openFamilyActivity(View view)
    {
        Intent i = new Intent(this, FamilyMembersActivity.class);
        startActivity(i);
    }
    public void openColorActivity(View view)
    {
        Intent i = new Intent(this, ColoursActivity.class);
        startActivity(i);
    }
    public void openPhraseActivity(View view)
    {
        Intent i = new Intent(this, PhrasesActivity.class);
        startActivity(i);
    }
}