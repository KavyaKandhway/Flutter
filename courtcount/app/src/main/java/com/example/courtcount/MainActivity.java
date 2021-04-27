package com.example.courtcount;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.view.View;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_main);
    }
    int score_A=0,score_B=0;

    public void inc3_A(View view)
    {
        score_A+=3;
        displayA(score_A);
    }
    public void inc2_A(View view)
    {
        score_A+=2;
        displayA(score_A);
    }
    public void inc1_A(View view)
    {
        score_A+=1;
        displayA(score_A);
    }
    public void inc3_B(View view)
    {
        score_B+=3;
        displayB(score_B);
    }
    public void inc2_B(View view)
    {
        score_B+=2;
        displayB(score_B);
    }
    public void inc1_B(View view)
    {
        score_B+=1;
        displayB(score_B);
    }
    public void displayA(int x)
    {
        TextView scoreView = (TextView) findViewById(R.id.scoreA);
        scoreView.setText(String.valueOf(x));
    }
    public void displayB(int x)
    {
        TextView scoreView = (TextView) findViewById(R.id.scoreB);
        scoreView.setText(String.valueOf(x));
    }
    public void resetAll(View view)
    {
        score_A=0;
        score_B=0;
        displayA(score_A);
        displayB(score_B);
    }
}