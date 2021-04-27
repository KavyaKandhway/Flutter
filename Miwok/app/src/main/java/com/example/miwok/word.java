package com.example.miwok;

public class word {
    private String miwok_word;
    private String eng_word;
    private int mImageResourceId;
    public word(String miwok,String eng,int ImageResourceId)
    {
        miwok_word=miwok;
        eng_word=eng;
        mImageResourceId=ImageResourceId;
    }
    public String getDefault()
    {
        return eng_word;
    }
    public String getMiwok()
    {
        return miwok_word;
    }
    public int getImageResourceId() {
        return mImageResourceId;
    }

}
