package com.example.miwok;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.ArrayAdapter;
import android.widget.LinearLayout;
import android.widget.ListView;
import android.widget.TextView;

import java.util.ArrayList;

public class NumbersActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_numbers);

        ArrayList<word> numbers=new ArrayList<word>();
        numbers.add(new word("lutti","one",R.drawable.number_one));
        numbers.add(new word("otiiko","two",R.drawable.number_two));
        numbers.add(new word("tolookosu","three",R.drawable.number_three));
        numbers.add(new word("oyyisa","four",R.drawable.number_four));
        numbers.add(new word("massokka","five",R.drawable.number_five));
        numbers.add(new word("temmokka","six",R.drawable.number_six));
        numbers.add(new word("kenekaku","seven",R.drawable.number_seven));
        numbers.add(new word("kawinta","eight",R.drawable.number_eight));
        numbers.add(new word("wo'e","nine",R.drawable.number_nine));
        numbers.add(new word("na'aacha","ten",R.drawable.number_ten));

        wordAdapter itemsAdapter = new wordAdapter(this,numbers,R.color.category_numbers);
        ListView listView = (ListView) findViewById(R.id.list_number);

        listView.setAdapter(itemsAdapter);



    }
}