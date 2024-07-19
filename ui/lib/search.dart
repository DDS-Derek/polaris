import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:infinite_scroll_pagination/infinite_scroll_pagination.dart';
import 'package:ui/providers/APIs.dart';
import 'package:ui/providers/settings.dart';
import 'package:ui/providers/welcome_data.dart';
import 'package:ui/widgets/progress_indicator.dart';

class SearchPage extends ConsumerStatefulWidget {
  const SearchPage({super.key, this.query});

  static const route = "/search";
  final String? query;

  @override
  ConsumerState<ConsumerStatefulWidget> createState() {
    return _SearchPageState();
  }
}

class _SearchPageState extends ConsumerState<SearchPage> {
  List<dynamic> list = List.empty();

  final PagingController<int, SearchResult> _pagingController =
      PagingController(firstPageKey: 1);
  int page = 1;
  @override
  void initState() {
    _pagingController.addPageRequestListener((pageKey) {
      _fetchPage(pageKey);
    });
    super.initState();
  }

  Future<void> _fetchPage(int pageKey) async {
    try {
      var searchList = await ref.read(
          searchPageDataProvider(Query(query: widget.query, page: pageKey))
              .future);
      print(searchList);
      final isLastPage = (searchList.page == searchList.totalPage);
      if (isLastPage) {
        _pagingController.appendLastPage(searchList.results!);
      } else {
        final nextPageKey = pageKey + 1;
        _pagingController.appendPage(searchList.results!, nextPageKey);
      }
    } catch (error) {
      _pagingController.error = error;
    }
  }

  @override
  void dispose() {
    _pagingController.dispose();
    super.dispose();
  }

  Widget _list(BuildContext context, List<SearchResult> data) {
    return NotificationListener(
      onNotification: (ScrollNotification scrollInfo) {
        if (scrollInfo is ScrollEndNotification &&
            scrollInfo.metrics.axisDirection == AxisDirection.down &&
            scrollInfo.metrics.pixels >= scrollInfo.metrics.maxScrollExtent) {
          viewModel.loadNextPage();
        }
        return true;
      },
      child: ListView.builder(
        itemCount: data.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text(data[index].name),
            subtitle: Text(data[index].description),
          );
        },
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    var list = PagedListView<int, SearchResult>(
      pagingController: _pagingController,
      builderDelegate: PagedChildBuilderDelegate<SearchResult>(
        itemBuilder: (context, item, index) {
          return Card(
              margin: const EdgeInsets.all(4),
              clipBehavior: Clip.hardEdge,
              child: InkWell(
                //splashColor: Colors.blue.withAlpha(30),
                onTap: () {
                  //showDialog(context: context, builder: builder)
                  _showSubmitDialog(context, item);
                },
                child: Row(
                  children: <Widget>[
                    Flexible(
                      child: SizedBox(
                        width: 150,
                        height: 200,
                        child: Image.network(
                          "${APIs.tmdbImgBaseUrl}${item.posterPath}",
                          fit: BoxFit.contain,
                          headers: APIs.authHeaders,
                        ),
                      ),
                    ),
                    Flexible(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Row(
                            children: [
                              Text(
                                "${item.name} ${item.name != item.originalName ? item.originalName : ''} (${item.firstAirDate?.year})",
                                style: const TextStyle(
                                    fontSize: 14, fontWeight: FontWeight.bold),
                              ),
                              const SizedBox(
                                width: 10,
                              ),
                              item.mediaType == "tv"
                                  ? const Chip(
                                      avatar: Icon(Icons.live_tv),
                                      label: Text(
                                        "电视剧",
                                      ))
                                  : const Chip(
                                      avatar: Icon(Icons.movie),
                                      label: Text("电影"))
                            ],
                          ),
                          const Text(""),
                          item.originCountry.isNotEmpty
                              ? Text("国家：${item.originCountry[0]}")
                              : Container(),
                          Text("${item.overview}")
                        ],
                      ),
                    )
                  ],
                ),
              ));
        },
      ),
    );

    return Column(
      children: [
        TextField(
          autofocus: true,
          controller: TextEditingController(text: widget.query),
          onSubmitted: (value) async {
            context.go(
                Uri(path: SearchPage.route, queryParameters: {'query': value})
                    .toString());
          },
          decoration: const InputDecoration(
              labelText: "搜索",
              hintText: "搜索剧集名称",
              prefixIcon: Icon(Icons.search)),
        ),
        Expanded(child: list)
      ],
    );
  }

  Future<void> _showSubmitDialog(BuildContext context, SearchResult item) {
    return showDialog<void>(
        context: context,
        builder: (BuildContext context) {
          return Consumer(
            builder: (context, ref, _) {
              String resSelected = "1080p";
              int storageSelected = 0;
              var storage = ref.watch(storageSettingProvider);

              return AlertDialog(
                title: Text('添加剧集: ${item.name}'),
                content: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    DropdownMenu(
                      label: const Text("清晰度"),
                      initialSelection: resSelected,
                      dropdownMenuEntries: const [
                        DropdownMenuEntry(value: "720p", label: "720p"),
                        DropdownMenuEntry(value: "1080p", label: "1080p"),
                        DropdownMenuEntry(value: "4k", label: "4k"),
                      ],
                      onSelected: (value) {
                        setState(() {
                          resSelected = value!;
                        });
                      },
                    ),
                    storage.when(
                        data: (v) {
                          return DropdownMenu(
                            label: const Text("存储位置"),
                            initialSelection: storageSelected,
                            dropdownMenuEntries: v
                                .map((s) => DropdownMenuEntry(
                                    label: s.name!, value: s.id))
                                .toList(),
                            onSelected: (value) {
                              setState(() {
                                storageSelected = value!;
                              });
                            },
                          );
                        },
                        error: (err, trace) => Text("$err"),
                        loading: () => const MyProgressIndicator()),
                  ],
                ),
                actions: <Widget>[
                  TextButton(
                    style: TextButton.styleFrom(
                      textStyle: Theme.of(context).textTheme.labelLarge,
                    ),
                    child: const Text('取消'),
                    onPressed: () {
                      Navigator.of(context).pop();
                    },
                  ),
                  TextButton(
                    style: TextButton.styleFrom(
                      textStyle: Theme.of(context).textTheme.labelLarge,
                    ),
                    child: const Text('确定'),
                    onPressed: () {
                      ref
                          .read(searchPageDataProvider(
                                  Query(query: widget.query, page: 1))
                              .notifier)
                          .submit2Watchlist(item.id!, storageSelected,
                              resSelected, item.mediaType!);
                      Navigator.of(context).pop();
                    },
                  ),
                ],
              );
            },
          );
        });
  }
}

class SearchBarApp extends StatefulWidget {
  const SearchBarApp({
    super.key,
    required this.onChanged,
  });

  final ValueChanged<String> onChanged;
  @override
  State<SearchBarApp> createState() => _SearchBarAppState();
}

class _SearchBarAppState extends State<SearchBarApp> {
  @override
  Widget build(BuildContext context) {
    return SearchAnchor(
        builder: (BuildContext context, SearchController controller) {
      return SearchBar(
        controller: controller,
        padding: const WidgetStatePropertyAll<EdgeInsets>(
            EdgeInsets.symmetric(horizontal: 16.0)),
        onSubmitted: (value) => {widget.onChanged(controller.text)},
        leading: const Icon(Icons.search),
      );
    }, suggestionsBuilder: (BuildContext context, SearchController controller) {
      return List<ListTile>.generate(0, (int index) {
        final String item = 'item $index';
        return ListTile(
          title: Text(item),
          onTap: () {
            setState(() {
              controller.closeView(item);
            });
          },
        );
      });
    });
  }
}
