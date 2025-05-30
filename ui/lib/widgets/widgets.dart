import 'package:flutter/material.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';
import 'package:ui/providers/APIs.dart';
import 'package:ui/widgets/utils.dart';

class Commons {
  static InputDecoration requiredTextFieldStyle({
    required String text,
    Widget? icon,
    String? helperText,
  }) {
    return InputDecoration(
      helperText: helperText,
      label: Row(
        children: [
          Text(text),
          const Text(
            "*",
            style: TextStyle(color: Colors.red),
          )
        ],
      ),
      icon: icon,
    );
  }
}

class SettingsCard extends StatelessWidget {
  final Widget? child;
  final GestureTapCallback? onTap;

  const SettingsCard({super.key, required this.child, this.onTap});

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.all(4),
      clipBehavior: Clip.hardEdge,
      child: InkWell(
          //splashColor: Colors.blue.withAlpha(30),
          onTap: onTap,
          child:
              SizedBox(width: 150, height: 150, child: Center(child: child))),
    );
  }
}

showLoadingWithFuture(Future f) {
  final context = APIs.navigatorKey.currentContext;
  if (context == null) {
    return;
  }
  showDialog(
    context: context,
    barrierDismissible: false, //点击遮罩不关闭对话框
    builder: (context) {
      return FutureBuilder(
          future: f,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.done) {
              if (snapshot.hasError) {
                return AlertDialog(
                  content: Text("处理失败：${snapshot.error}"),
                  actions: [
                    TextButton(
                        onPressed: () => Navigator.of(context).pop(),
                        child: const Text("好"))
                  ],
                );
              }
              Navigator.of(context).pop();
              return Container();
            } else {
              return const AlertDialog(
                content: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: <Widget>[
                    CircularProgressIndicator(),
                    Padding(
                      padding: EdgeInsets.only(top: 26.0),
                      child: Text("正在处理，请稍后..."),
                    )
                  ],
                ),
              );
            }
          });
    },
  );
}

class MyRangeSlider extends StatefulWidget {
  final String name;
  const MyRangeSlider({super.key, required this.name});

  @override
  State<StatefulWidget> createState() {
    return _MySliderState();
  }
}

class _MySliderState extends State<MyRangeSlider> {
  double sizeMax = 5000;
  @override
  Widget build(BuildContext context) {
    return FormBuilderRangeSlider(
        decoration: const InputDecoration(labelText: "文件大小限制"),
        maxValueWidget: (max) => Text("${sizeMax / 1000} GB"),
        minValueWidget: (min) => const Text("0"),
        valueWidget: (value) {
          final sss = value.split(" ");
          return Text("${readableSize(sss[0])} - ${readableSize(sss[2])}");
        },
        onChangeEnd: (value) {
          if (value.end > sizeMax * 0.9) {
            setState(
              () {
                sizeMax = sizeMax * 5;
              },
            );
          } else if (value.end < sizeMax * 0.2) {
            if (sizeMax > 5000) {
              setState(
                () {
                  sizeMax = sizeMax / 5;
                },
              );
            }
          }
        },
        name: widget.name,
        min: 0,
        max: sizeMax);
  }

  String readableSize(String v) {
    if (v.endsWith("K")) {
      return v.replaceAll("K", " GB");
    }
    return "$v MB";
  }
}

class LoadingListTile extends StatefulWidget {
  const LoadingListTile(
      {super.key,
      required this.onPressed,
      required this.icon,
      required this.text});
  final Future<void> Function() onPressed;
  final IconData icon;
  final String text;

  @override
  State<StatefulWidget> createState() {
    return LoadingListTileState();
  }
}

class LoadingListTileState extends State<LoadingListTile> {
  bool loading = false;

  @override
  Widget build(BuildContext context) {
    return ListTile(
        onTap: loading
            ? null
            : () async {
                setState(() => loading = true);
                try {
                  await widget.onPressed();
                } catch (e) {
                  showSnakeBar("操作失败：$e");
                } finally {
                  setState(() => loading = false);
                }
              },
        title: Text(widget.text),
        leading: loading
            ? Container(
                width: 24,
                height: 24,
                padding: const EdgeInsets.all(2.0),
                child: const CircularProgressIndicator(
                  color: Colors.grey,
                  strokeWidth: 3,
                ),
              )
            : Icon(widget.icon));
  }
}

class LoadingIconButton extends StatefulWidget {
  const LoadingIconButton(
      {super.key, required this.onPressed, required this.icon, this.tooltip});
  final Future<void> Function() onPressed;
  final IconData icon;
  final String? tooltip;

  @override
  State<StatefulWidget> createState() {
    return _LoadingIconButtonState();
  }
}

class _LoadingIconButtonState extends State<LoadingIconButton> {
  bool loading = false;

  @override
  Widget build(BuildContext context) {
    return IconButton(
        tooltip: widget.tooltip,
        onPressed: loading
            ? null
            : () async {
                setState(() => loading = true);
                try {
                  await widget.onPressed();
                } catch (e) {
                  showSnakeBar("操作失败：$e");
                } finally {
                  setState(() => loading = false);
                }
              },
        icon: loading
            ? Container(
                width: 24,
                height: 24,
                padding: const EdgeInsets.all(2.0),
                child: const CircularProgressIndicator(
                  color: Colors.grey,
                  strokeWidth: 3,
                ),
              )
            : Icon(widget.icon));
  }
}

class LoadingTextButton extends StatefulWidget {
  const LoadingTextButton(
      {super.key, required this.onPressed, required this.label});
  final Future<void> Function() onPressed;
  final Widget label;

  @override
  State<StatefulWidget> createState() {
    return _LoadingTextButtonState();
  }
}

class _LoadingTextButtonState extends State<LoadingTextButton> {
  bool loading = false;

  @override
  Widget build(BuildContext context) {
    return TextButton.icon(
      onPressed: loading
          ? null
          : () async {
              setState(() => loading = true);
              try {
                await widget.onPressed();
              } catch (e) {
                showSnakeBar("操作失败：$e");
              } finally {
                setState(() => loading = false);
              }
            },
      icon: loading
          ? Container(
              width: 24,
              height: 24,
              padding: const EdgeInsets.all(2.0),
              child: const CircularProgressIndicator(
                color: Colors.grey,
                strokeWidth: 3,
              ),
            )
          : Text(""),
      label: widget.label,
    );
  }
}

class LoadingElevatedButton extends StatefulWidget {
  const LoadingElevatedButton(
      {super.key, required this.onPressed, required this.label});
  final Future<void> Function() onPressed;
  final Widget label;

  @override
  State<StatefulWidget> createState() {
    return _LoadingElevatedButtonState();
  }
}

class _LoadingElevatedButtonState extends State<LoadingElevatedButton> {
  bool loading = false;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton.icon(
      onPressed: loading
          ? null
          : () async {
              setState(() => loading = true);
              try {
                await widget.onPressed();
              } catch (e) {
                showSnakeBar("操作失败：$e");
              } finally {
                setState(() => loading = false);
              }
            },
      icon: loading
          ? Container(
              width: 24,
              height: 24,
              padding: const EdgeInsets.all(2.0),
              child: const CircularProgressIndicator(
                color: Colors.grey,
                strokeWidth: 3,
              ),
            )
          : Text(""),
      label: widget.label,
    );
  }
}

class PoError extends StatelessWidget {
  const PoError({super.key, required this.msg, required this.err});
  final String msg;
  final dynamic err;

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          "$msg ",
          style: TextStyle(color: Theme.of(context).colorScheme.error),
        ),
        Tooltip(
          message: "$err",
          child: Icon(
            Icons.info,
            color: Theme.of(context).colorScheme.error,
          ),
        )
      ],
    );
  }
}

class PoNetworkError extends StatelessWidget {
  const PoNetworkError({super.key, required this.err});
  final dynamic err;
  @override
  Widget build(BuildContext context) {
    return PoError(msg: "网络错误，请检查网络链接", err: err);
  }
}

class PoProgressIndicator extends StatelessWidget {
  const PoProgressIndicator({super.key, this.backgroundColor, this.value, this.icon});
  final double? value;
  final Color? backgroundColor;
  final IconData? icon;

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: AlignmentDirectional.center,
      children: [
        CircularProgressIndicator(
          backgroundColor: backgroundColor,
          value: value,
        ),
        icon != null ? Opacity(opacity: 0.7, child: Icon(icon, color: Theme.of(context).colorScheme.primary,),):Container()
      ],
    );
  }
}
