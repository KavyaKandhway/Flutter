import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter/services.dart';
import 'package:flutter_stepper_app/label_widget.dart';
import 'package:im_stepper/stepper.dart';
import 'package:pin_entry_text_field/pin_entry_text_field.dart';

void main() {
  runApp(DotStepperDemo());
}

final name = TextEditingController();
final email = TextEditingController();
final phoneCode = TextEditingController(text: '+91');
final number = TextEditingController();

class DotStepperDemo extends StatefulWidget {
  @override
  _DotStepperDemo createState() => _DotStepperDemo();
}

class _DotStepperDemo extends State<DotStepperDemo> {
  int activeStep = 1;
  int lowerBound = 0;
  int upperBound = 0;
  String label = 'Human';
  List<Widget> elements = [
    Column(
      children: [
        Container(
          padding: const EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 0.0),
          child: Text('What do we call you?',
              textAlign: TextAlign.left,
              style: TextStyle(
                color: Color(0xFF5C5C77),
                fontSize: 22.0,
                fontWeight: FontWeight.w500,
              )),
          height: 40.0,
          width: 350.0,
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 0.0, 25.0, 20.0),
          child: TextField(
            controller: name,
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.zero,
                borderSide: BorderSide(color: Color(0xFFF7A11C)),
              ),
              labelText: 'Name',
              //hintText: 'Enter Your Name'
            ),
          ),
        ),
        Container(
          padding: const EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 0.0),
          child: Text('Your Email?',
              textAlign: TextAlign.left,
              style: TextStyle(
                color: Color(0xFF5C5C77),
                fontSize: 22.0,
                fontWeight: FontWeight.w500,
              )),
          height: 40.0,
          width: 350.0,
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 0.0, 25.0, 10.0),
          child: TextField(
            controller: email,
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.zero,
                borderSide: BorderSide(color: Color(0xFFF7A11C)),
              ),
              labelText: 'Email',
              //hintText: 'Enter Your Name'
            ),
          ),
        ),
      ],
    ),
    Column(
      children: [
        Container(
          padding: const EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 10.0),
          child: Text('Can we get your number :P',
              textAlign: TextAlign.left,
              style: TextStyle(
                color: Color(0xFF5C5C77),
                fontSize: 22.0,
                fontWeight: FontWeight.w500,
              )),
          height: 40.0,
          width: 350.0,
        ),
        Row(
          mainAxisSize: MainAxisSize.max,
          children: [
            Expanded(
              flex: 3,
              child: Padding(
                padding: const EdgeInsets.fromLTRB(25.0, 0.0, 5.0, 0.0),
                child: TextField(
                  maxLength: 3,
                  controller: phoneCode,
                  decoration: InputDecoration(
                    counterText: "",
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.zero,
                      borderSide: BorderSide(color: Color(0xFFF7A11C)),
                    ),
                  ),
                ),
              ),
            ),
            Expanded(
              flex: 10,
              child: Padding(
                padding: const EdgeInsets.fromLTRB(2.0, 0.0, 25.0, 0.0),
                child: TextField(
                  keyboardType: TextInputType.numberWithOptions(decimal: true),
                  inputFormatters: [FilteringTextInputFormatter.digitsOnly],
                  maxLength: 10,
                  controller: number,
                  decoration: InputDecoration(
                    counterText: "",
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.zero,
                      borderSide: BorderSide(color: Color(0xFFF7A11C)),
                    ),
                    labelText: 'Phone Number',
                    //hintText: 'Enter Your Name'
                  ),
                ),
              ),
            ),
          ],
        ),
        Container(
          padding: const EdgeInsets.fromLTRB(10.0, 25.0, 0.0, 0.0),
          child: Text('OTP',
              textAlign: TextAlign.left,
              style: TextStyle(
                color: Color(0xFF5C5C77),
                fontSize: 22.0,
                fontWeight: FontWeight.w500,
              )),
          height: 50.0,
          width: 350.0,
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 10.0, 25.0, 0.0),
          child: PinEntryTextField(
            showFieldAsBox: true,
          ),
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 10.0, 10.0, 0.0),
          child: Label(
              title: 'Did not receive? Send again.',
              labelColor: 0xFF5C5C77,
              labelSize: 15.0),
        )
      ],
    ),
    Column(
      children: [
        Container(
          padding: const EdgeInsets.fromLTRB(10.0, 10.0, 0.0, 0.0),
          child: Text('Where should we deliver treats?',
              textAlign: TextAlign.left,
              style: TextStyle(
                color: Color(0xFF5C5C77),
                fontSize: 22.0,
                fontWeight: FontWeight.w500,
              )),
          height: 40.0,
          width: 350.0,
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 5.0, 25.0, 10.0),
          child: Row(
            children: [
              Expanded(
                flex: 1,
                child: Padding(
                  padding: const EdgeInsets.only(right: 8.0),
                  child: TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.zero,
                        borderSide: BorderSide(color: Color(0xFFF7A11C)),
                      ),
                      labelText: 'Flat No.',
                    ),
                  ),
                ),
              ),
              Expanded(
                flex: 1,
                child: TextField(
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.zero,
                      borderSide: BorderSide(color: Color(0xFFF7A11C)),
                    ),
                    labelText: 'Pin Code',
                  ),
                ),
              ),
            ],
          ),
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 0.0, 25.0, 10.0),
          child: TextField(
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.zero,
                borderSide: BorderSide(color: Color(0xFFF7A11C)),
              ),
              labelText: 'Address Line 1',
            ),
          ),
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(25.0, 0.0, 25.0, 0.0),
          child: TextField(
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.zero,
                borderSide: BorderSide(color: Color(0xFFF7A11C)),
              ),
              labelText: 'Address Line 2',
            ),
          ),
        ),
      ],
    ),
  ];

  @override
  Widget build(BuildContext context) {
    Future<bool> onBack() {
      if (activeStep > lowerBound) {
        setState(() {
          activeStep--;
        });
      } else {
        Navigator.pop(context);
      }
    }

    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: SafeArea(
        child: Scaffold(
          body: Column(
            mainAxisAlignment: MainAxisAlignment.end,
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisSize: MainAxisSize.max,
            children: [
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: Container(
                  child: WillPopScope(
                    onWillPop: onBack,
                    child: DotStepper(
                      dotCount: 3,
                      activeStep: activeStep,
                      lowerBound: (bound) => lowerBound = bound,
                      upperBound: (bound) => upperBound = bound,
                    ),
                  ),
                ),
              ),
              SizedBox(
                height: 30.0,
              ),
              Container(
                padding: const EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 0.0),
                child: Text('Hello, $label',
                    textAlign: TextAlign.left,
                    style: TextStyle(
                      color: Color(0xFF403B87),
                      fontSize: 35.0,
                      fontWeight: FontWeight.w600,
                    )),
                height: 50.0,
                width: 350.0,
              ),
              Container(
                padding: const EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 0.0),
                child: Text('Let\'s get you onboard',
                    textAlign: TextAlign.left,
                    style: TextStyle(
                      color: Color(0xFFF7A11C),
                      fontSize: 18.0,
                      fontWeight: FontWeight.w400,
                    )),
                height: 40.0,
                width: 350.0,
              ),
              elements[activeStep - 1],
              Expanded(
                child: SizedBox(
                  height: 300.0,
                ),
              ),
              Container(
                child: nextButton(),
              )
            ],
          ),
          resizeToAvoidBottomPadding: false,
        ),
      ),
    );
  }

  /// Returns the next button.
  Widget nextButton() {
    return FlatButton(
      onPressed: () {
        if (activeStep < upperBound) {
          setState(() {
            activeStep++;
          });
        }
      },
      minWidth: 1000,
      height: 60.0,
      child: Text(
        'Next',
        style: TextStyle(
          fontSize: 15.0,
          fontWeight: FontWeight.w900,
        ),
      ),
      color: Color(0xFF403B87),
      textColor: Colors.white,
    );
  }

  /// Returns the previous button.
  Widget previousButton() {
    return IconButton(
        icon: new Icon(Icons.arrow_back, color: Color(0xFF403B87)),
        onPressed: () {
          if (activeStep > lowerBound) {
            setState(() {
              activeStep--;
            });
          }
        });
  }
}
