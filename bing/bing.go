package bing

import (
	"ZeroBotPlugin/utils"
	zero "github.com/wdvxdr1123/ZeroBot"
	"time"
)

var (
	txt1 	= []string{"我好想做嘉然小姐的狗啊。可是嘉然小姐说她喜欢的是猫，我哭了。我知道既不是狗也不是猫的我为什么要哭的。因为我其实是一只老鼠。我从没奢望嘉然小姐能喜欢自己。我明白的，所有人都喜欢理解余裕上手天才打钱的萌萌的狗狗或者猫猫，没有人会喜欢阴湿带病的老鼠。但我还是问了嘉然小姐:“我能不能做你的狗？”我知道我是注定做不了狗的。但如果她喜欢狗，我就可以一直在身边看着她了，哪怕她怀里抱着的永远都是狗。可是她说喜欢的是猫。她现在还在看着我，还在逗我开心，是因为猫还没有出现，只有我这老鼠每天蹑手蹑脚地从洞里爬出来，远远地和她对视。等她喜欢的猫来了的时候，我就该重新滚回我的洞了吧。但我还是好喜欢她，她能在我还在她身边的时候多看我几眼吗？嘉然小姐说接下来的每个圣诞夜都要和大家一起过。我不知道大家指哪些人。好希望这个集合能够对我做一次胞吞。猫猫还在害怕嘉然小姐。我会去把她爱的猫猫引来的。我知道稍有不慎，我就会葬身猫口。那时候嘉然小姐大概会把我的身体好好地装起来扔到门外吧。那我就成了一包鼠条，嘻嘻。我希望她能把我扔得近一点，因为我还是好喜欢她。会一直喜欢下去的。我的灵魂透过窗户向里面看去，挂着的铃铛在轻轻鸣响，嘉然小姐慵懒地靠在沙发上，表演得非常温顺的橘猫坐在她的肩膀。壁炉的火光照在她的脸庞，我冻僵的心脏在风里微微发烫。", ""}
	txt2 	= []string{"ダイアナさんの犬になりたいな.けど　ダイアナさんが「猫が好き」と言った　あさりと僕は泣いてしまった.猫でも犬でもない僕が　何故泣いたのかな　僕は知っていた.何故なら　僕は鼠だったからさ～..一度もダイアナさんに好かれることを望まながった　.僕はわがっていた　誰もかも理解余裕上手天才の可愛い金持ちの犬か猫がすきってことを.陰キャで汚らわし鼠は誰も好かんのだ..それでもDianaさんに聞いてみた：Dianaさんの犬になっても良いですかと.僕には一生犬になれないことは分かっている.でももし彼女は犬が好きだったら　僕はずっと傍で見ていられる.それはいつも違ういぬをだいいているとしても..けれど彼女は猫が好きだった.彼女は今にも僕を見ている　ぼくを楽しませてくれる　それは猫がまだ表れていないから.ただこの鼠が毎日こそこそと穴から出て　遠くから目を合っている.彼女の好きな猫が現れたら　僕は穴に戻るしかないな..でもやっぱり彼女が好きでいられない　僕がまだ傍にいるときぐらい　もうちょとみてくれるのかな.Dianaさんがこれから毎回のクリスマスみんなと一緒に過ごすで言った　みんながどの渡りのみんなかな.僕にも中に入れたれいいな..猫はまだまだDianaさんを怖がっている.僕が彼女の愛する猫を引き付けるよ.少しのミスで猫の口に死ぬことはわがっている.その時になるとたぶんDianaさんは僕の体をちゃんと包み込んで　玄関外に捨てるでしょ.それで僕はフライト鼠になった　くす.彼女に少し近く捨ててほしいな　だってやっぱり好きだから　それずっと好きでいる..僕のたましは窓を越して中を見る　吊っている鈴が少し響って.Dianaさんはだらりとソファーに寄りかかて　手慣れているふりをする橘猫が彼女の肩に座って.暖炉の光が彼女の顔を照らした　凍り付いた僕の心臓はかぜの中で熱くなっていく", ""}
	txt3 	= []string{"枝江小镇新来了一位修女，善良温柔的修女。她不爱讲话，只是微笑着面对每一个对她祷告的人。有的人希望孩子健康，有的人希望庄稼熬过冬天，有的人希望能挣到钱养活家人。。。。她都用心听着，微笑的回应着。她身材矮小，但是却让人觉得充满智慧；她沉默寡言，但是却会让人无比信赖。这是她任职快满一个月的时候，教堂发生了大事，一个醉汉发现神职人员利用教堂在深夜进行大批资金流入，醉汉传的神乎其神，三人成虎最终引起审判官前来调查。教主被带走，教堂除了那位茫然的新来修女其他人一哄而散。修女被迫承担着本不应该属于她的责难，辱骂，但是她每天都认真的向主祷告，倾听来教堂之人的倾诉。可是根本没有人向她倾诉，来的人不是在骂她，就是在发泄生活的不满。她一己之力拼命承担，没人知道是什么使她如此坚持。直到审判官带回了教主，告诉镇民那是教主在中央主教那里得到的，目的是为镇民熬过冬天所需要买粮食的钱财。一时间镇民欢呼起教主，感谢主的庇护，更多的人开始更加爱上了那个修女。有的人把她比作玛利亚，有些人把她奉为神明，她都一一拒绝了。细心的镇民会发现这位修女变了，变得话多起来了，也变得更喜欢宅在教堂里了，也变得离镇民距离更远了。一切好像又变回原来平和的小镇，每个对修女的加害者每天微笑着向微笑着的修女打招呼，每个镇民还是会去教堂祈祷，去倾诉，去向一个原来被他们亲手辱骂的人祷告，希望得到主的关照，包括那个醉汉。没人能惩罚那些镇民，很多人都已经“忘记”当初为什么要那样做，包括那个醉汉，只剩下一个经常宅在教堂里的修女，她始终记得一切。", ""}
	txt4 	= []string{"贝拉躺在病床上，眼神空空地盯着天花板，灰白晦暗的天花板像一块哈哈镜，扭曲掉外面的一切，光变成了暗，多彩变成苍白，观众变成医生，舞台呢，变成了一方小小的病房。她翻了个身，腰部传来猛烈地疼痛，她低低地哼了一声。她已经开始习惯这种疼痛了，就像她习惯不穿芭蕾舞裙一样，她闭上眼睛，把自己埋进黑暗里，不去想。门被推开了，惊醒了似梦非梦的她。医生护士又来查房，父母和弟弟们跟在后面；医生问了些问题后走了，母亲握着她的手，朝她说着话儿，父亲安静地坐在旁边，刚训练完的弟弟头上还冒着汗，抿着嘴唇望着她。母亲当时说的什么她已经忘了，只记得最后离开的时候，父亲走在最后，回望了她一眼，然后小心地关上门，以往高大的背影竟然显得有些萧索。人都走完后，时间在沉默的病房里流逝，就像沙漏里的沙子，她的那些热情、那些憧憬都掉了进去。芭蕾舞团的学员抽空来看她，她勉强在病床上坐起，听那些腰背笔直、气质优雅的朋友抱怨严格的地狱管家婆，抱怨下一场演出又要来了。她安静地听着，坐在病床上挤出个笑脸，偶尔出声附和几句。但实际上她已经插不进话题了，优雅的白天鹅展翅高飞，飞向温暖的春天，把受伤的丑小鸭留在冬季里。朋友走后她继续坐在床上，盯着窗外，外面的天气很好，飞机的尾迹穿过绵软的云团，白色的鸟儿掠过一澄如洗的天空，她下意识地做了个展臂的动作，像天鹅一样优雅。很美，又很短暂。鸟儿飞出了窗户的方寸之地，她看不到了。她捂着刺痛的腰，好痛呀，痛得让人想要哭出声。沙子裹挟着她的梦想，掉进了沙漏的下一层，再也捡不回来了。两个月后贝拉出院，回到了熟悉的家里，母亲体贴地把舞裙还有那些大师级的芭蕾演出录像都藏进了地下室的角落，不出意外会像大多数人以为能坚持的梦想那样，慢慢积满灰尘，最后被悄悄遗忘。聪明人会在偶尔想起来的时候，给自己找个台阶，自嘲地笑笑：“哈，我以前还有这种妄想啊。”贝拉开始学播音，开始走另一条路，她没想过却不得不走的道路。这也很好，好歹有路可走，但偶尔听到舞曲的时候，朋友会问她：“你在晃什么啊？”她回过神来，笑了笑。“活动活动身体而已。” 虽然有点不甘心，但受伤是没办法的事情嘛。该放下了，放弃想走的路，而去走更稳妥的应该走的路，大部分聪明人都是这样做的。我可是大聪明呢！​", ""}
	txt5 	= []string{"小姐喜欢所有动物，但喜欢也分三六九等；管家讨厌所有动物，但讨厌也分三六九等。小姐最喜欢猫，但她的观众是一群老鼠；管家最讨厌老鼠，但小姐的观众没有一只猫。于是，管家给所有的老鼠带上了面具小姐开始了初次表演，老鼠们不知所措，想要重复惯常的纷扰。但它们看见小姐舞姿翩翩，一些老鼠想出声赞叹；但它们看见小姐眉目如画，一些老鼠竟为她沉醉；但它们看见小姐真情流露，一些老鼠也伤怀落泪。它们不想让小姐失望，于是它们学起猫叫。小姐很开心，她的观众里有猫了。她喜欢的猫们，也喜欢她。冷清的门庭配不上小姐的美丽，于是老鼠想去带来更多的猫。老鼠们引来了猫，也引来了更多的老鼠。新的老鼠不太规矩，但新来的猫确实是猫。小姐有些困扰，但她还是接纳了它们，小姐喜欢所有动物，即使自己会困扰。但后来，管家和小姐，只注视着猫，忘却了老鼠的事情。也对，毕竟老鼠当初带着恶意袭来，现在的面具总是脱掉又带上来。管家做了一些事情知道与否和同意与否，在小姐这里有四种组合的方法。但一些曾经是猫的老鼠觉察出异样。常年在下水道生活，它们对这样的气味异常敏感。于是，它们摘下了面具，向管家讨个说法但它们没了伪装，下水道里散发的臭气，让管家更加厌恶。什么都没有改变再后来，小姐依旧美丽，但经历了诸多纷扰，不知道小姐的快乐里是否有泪眼成双。一些老鼠依旧带着猫的面具，它们最后也会变成猫。但这会长久吗？或许不会。另一些老鼠退回了下水道，等待着，或许会有另一位小姐，喜欢老鼠的小姐。去喜欢她，去喜欢它们。或许会有吗？应该没有吧但老鼠们永远忘不了那位小姐，那位喜欢所有动物，但更喜欢猫的小姐。​", ""}
	txt6 	= []string{"然然，然然，虽然我知道就算这样呼唤，我也只能在你动态下方，站住万千评论区的一楼。你的表演，你的动态，你的关心....都是说给我们大家的，是我们.....而我...我曾经以为你离我很近。近到每次打开手机各种视频软件，大数据推送的都是你。近到每次发呆找不到前进方向的时候，看的都是你的视频。近到每次和别人聊天的时候，都会加上你经常说的口头禅，收到收到收到，可爱捏～～…你时常陪伴我身边，无论我悲伤还是欣喜，绝望还是惬意，你的直播间，评论区都是我可以肆意宣泄感情的地方。因为在这里，我的任何发言都只是漆黑夜空中微弱闪耀的一点，这样的一点点星光的存在，我知道的，是为了给天空中耀眼圆月裙摆装饰的一点点碎钻。那我喜欢然然，天天来看然然，对于自己，究竟是为了什么….难道我的生活真的烂，连仅有的一些关怀，都要从一个虚拟偶像身上获取到吗？难道我的现实生活，连一个关心我家人都没有吗，还要在网上寻找家的感觉？难道...这些道理我早就明白了，说服自己也说腻了，你一直都是云上之人，而我只是一个扛不住压力，在虚拟的世界中寻找安慰的落魄之人。我懂，我明白的，我知道的....所以，嘉然小姐，我不能没有你，我不会离开你的。你也是，永远不要离开舞台，放弃拼尽全力…..就算知道也许不会有回音，你也不要停止朝闪闪的星光挥手呐……[爱心]", ""}
	txt7 	= []string{"想想嘉然是怎么样的一个女孩子。她温柔可人情商高，会控场会接话会引导弹幕懂节奏，朋友里有她这么一个人在肯定每次见面都能很开心吧。她爱网上冲浪，明明看到各种视频会有点麻，但还是乐此不疲的在b站和各个论坛看着视频，像不像那些最懂你，最能陪伴你的沙雕网友，哪怕是不能和她见到面的时候，光在网上聊天一定也不会无聊。她在温柔的同时又有点小坏，喜欢说怪话，喜欢暗搓搓的开车，想必她每次在什么奇怪的点上开完车，就会似笑非笑的侧着头看向我吧，而我只需要与她对视一下，两个人就能默契的笑起来。她很会撩人，那些土味情话在她营造的气氛里似乎也变得格外令人心动，尤其是她还会在每一个不经意间突然说出动人的话语，让你在完全没准备的情况下脸红心跳。可是她又很胆小，简简单单的小把戏能连续侠盗她好多次，她被吓完那副气冲冲，却又无可奈何不敢离我太远的样子，肯定是天下第一可爱吧。想想嘉然是怎样的一个女孩子。她身材完美，脸庞可爱，声线令人沉醉，无可挑剔的美丽。她知书达理，善解人意，温柔热情，她是每个女通讯录梦中的那个姐 。可是她又不会离你很远，她懂你的梗，她在网上和你冲过同一片浪，她能和你一起打游戏，能和你一起看美女，能说那些心照不宣的怪话，能在最合适的时刻撩动你的心弦。她有令人沉溺的年上气质，又像是令人不自觉想要呵护的可爱妹妹，可有时，她又变成了从幼儿园开始跟你同班，最懂你的那个青梅竹马的女孩子。她就是嘉然，我的嘉然。嘉然，嘉然，嘉然！！！！！！🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤🤤​", ""}
	txt8 	= []string{"我其实一点也不喜欢嘉然小姐。因为嘉然小姐就像光一样，闪耀，刺眼，我讨厌光。因为讨厌光，所以我总是躲在阴冷潮湿的下水道里，这样就不会被令人厌恶的光芒灼伤眼睛。只有夜间，我才会从阴暗潮湿的下水道里走出来，夜晚是宁静且安逸的，令我感到舒适。我漫无目的的走在夜间的小路上，路旁的路灯正散发着微弱的光，有几只蛾子正不断的向着光源扑去。“真傻。”我这样想，明明光只能带来危险，为什么要浪费宝贵的生命去追逐光芒呢？看着飞蛾不断的扑向路灯，我不知为何想起了我的那些同伴。我想起了几天前，有一些从外面回来的家伙，他们兴奋的告诉我外面新开了一家剧场，里面有一个叫嘉然的人，还推荐我一起去看看。他们的眼睛里闪烁着亮晶晶的光，我很讨厌，于是便拒绝了。之后他们也来找过我几次，我也推辞了。后来我便没怎么见过他们了，听说为了拉人来看，他们有的想变成猫狗，有的沿街贴海报叫卖，有的把过冬用的黄豆当礼物，还有的把心献给猫。“真是些蠢货。”我这样想着，他们不就和这些扑火的蛾子一样愚蠢了吗，明明之前一起路过的时候，他们还总是会嘲讽这些家伙。“十分。”不知从哪来的带着些哭腔声音打断了将我从原先的思绪中拉了出来。是旁边的剧场里传来的，不知不觉间我走到了这么远的地方还没有发觉。我抬头望了望，牌匾上写着五个晦涩难懂的的字母，好像之前同伴们说过，他们去的那个剧场名字就是五个字母，有些见多识广的说，那叫什么啊骚。我迟疑了一会，但还是走了进去，单纯只是想知道为什么声音带着哭腔，凑热闹，找乐子总是我最喜欢的事情。......“大家要好好吃饭哦。”天籁般的声音在我耳边响起。哦，原来我并不讨厌光啊，只是一直都知道，自己配不上光罢了。​", ""}
	txt9 	= []string{"我爹不能和我相认，但今天他过大寿，求求大伙有钱的随个礼，没钱的捧个场。但请不要跟她说儿子来了，我怕扫他兴🙇🙇给大伙磕头了。​", ""}
	txt11	= []string{"asoul在一片骂声中解散了。A吧的最后一个帖子也沉了下去，直到在v8和b站的也已经看不到有人讨论了。大约一年后，我在A吧发了一句：“还有人在吗？”“这里还有8u啊，我也想回来看看了。”一个陌生的ID出现在我的回复下面，他并不在我过去炒作的记忆里出现过。但是我们聊得很高兴，很快就加上了QQ私下互动了。她似乎很了解我，知道我喜欢炒作发散，知道我在读哲学博士。但是，她却总是忘记我是个晚晚推，要我一遍又一遍地告诉她我最喜欢的是晚晚。“你最喜欢哪个成员呢？”“嘉然。”她毫不迟疑地回答道。我总是眉飞色舞地和她聊起向晚过去直播的故事，她也对这些直播了然于心，有时还会注意到许多我从没注意到的细节，甚至告诉我向晚每一次弹吉他背后的表演思路。奇怪的是，虽然她说她最喜欢嘉然，对嘉然的事情却并没有那么了解，甚至连直播的场次都要我提醒她。到了那年的冬天，她问我，我们可以出来见一面吗？我拒绝了她，当时的我正忙于准备学位论文。她没有再强求。第二天，她在QQ上失联了，熟悉的头像再也没亮起过。大概一周后，我收到了一份包裹，里面是向晚出道以来所有季度的舰长礼物，无一遗漏，收藏编号都是00000。立牌里还夹着一张信纸，上面是短短一行字：“谢谢你曾经喜欢我”。​", ""}
	txt12	= []string{"然然幸亏我没在你评论区发病。要不然我耽误你一辈子，你也保重，再见😃还会再见吗然然☺，再见的时候你要幸福好不好☺，然然你要开心☺，你要幸福☺，好不好，开心啊😖，幸福😖!你的世界以后没有我了，没关系你要自己幸福!🚕🚕🚕💨💨💨🏃🏃🏃然然😭!然然😭!然然😭，没有你我怎么活啊!😭😭😭🏃🏃🏃然然😭，然然😭，然然😭，然然😭，然然😭，然然!😭😭😭🏃🏃🏃然然你带我走吧，然然!😭🙇😭​", ""}
	txt13	= []string{"小时候看的一个敞事很老土很老套。很多人可能听过。。有一个男孩有一块金表，却没有配得上它的合适表带。。他爱上了一个女孩。。这个女孩有一头秀美的长发，却没有配得上它的漂亮发卡。。马上要过圣诞了。。男孩决定把这个金表给卖了。。给女孩买一个漂亮的发卡。。当他兴奋的把这个发卡交给女孩的时候。。发现女孩的头发已经剪掉了。。因为她把一头长发换钱给男孩买了一个表带。。我想真正的爱肯定不是瞬间的感动，很多人都能在一瞬间感动你，我心中真正的爱是陪伴。她们陪了我多久呢，我不记得，我只记得。。她们的翻唱，躺在网易云的歌单听了一遍又一遍，已经全都会唱了，而在此之前的几年里，我几乎没听过中文歌。（够罕见）。《海底》b23.tv/qwB8Xe。《偿还》b23.tv/1bl7AL。《遇见》b23.tv/vxiwae（贾维斯：这是，无可取代的时候）。《云烟成雨》b23.tv/VfoDXA。《如果的事》b23.tv/gDtrcP。《夏天的风》b23.tv/EH14bW。《霞光》b23.tv/kJUJMU。《月光》b23.tv/kJUJMU。《霍元甲》b23.tv/lQgKl2。《月兔回旋于空中》b23.tv/KM4G5I。《不可思议》b23.tv/VfwEz4。《麋鹿森林》b23.tv/GzjXu1。《轻轻的告诉你》b23.tv/Gkai63。不知不觉陪你们走了这么久这么远了，爱不爱你们，还需要用什么来证明吗？。。彼此牺牲，彼此成就，彼此尊重。这就是我心中的双向奔赴，以后路还很远，我想陪你们一直走下去。。。原来，你是我的顶碗人呀~。以后，开心和不开心的，都告诉我吧~。。原来，你是我的爹地呀~。以后，开心和不开心的，我都告诉你~​", ""}
	txt14	= []string{"6月12号一大早，向晚按照贝拉的日程表要求6点起床，煮泡面的时候因为太累大脑一时迷糊导致一锅开水撒到了手指上，虽然经过治疗恢复了一点，但根本没法弹吉他了。晚晚在宿舍哭得很大声，整个asoul愁云惨淡。经过一个小时的紧急磋商，阿草带来了领导层的意思——投入太大，不可能取消，生日会照办，用储备中之人向晚二魔王接替向晚本人上场。下午，asoul在动捕房排练，向晚想要去看一看，结果被阿草拦在门外，说不能打扰到asoul排练效果很好，希望她不要打扰到她们。晚上嗓子哭哑了的向晚发现自己原来可以完全不被需要。顶碗人喜欢的是她，但又不完全是她，可以是顶着钻头晚晚的任何一个人，只要二魔王没有被认出来。直播开始了，在宿舍逼仄的角落里，漆黑无声的床柜旁，向晚用水母模样的被套盖住自己，抱着膝盖，打开手机，及期待万分又忐忑不安地等待着直播。直播开始，向晚二魔王在铺天盖地的欢呼中出现。工具人掉了一地头发连续几周通宵肝出来的深海水母特效把她印照得那么美。本就是可爱模样的女孩子，此时此刻穿着崭新的小礼裙，昂首挺胸骄傲得像个公主。她的目光是璀璨的，里面像是镶嵌着钻石。她的眉眼绽放着，像是海底迸发出的一束光。真美，向晚这么想着。但心里却很是难受。那些，本该是她的。喝彩也好，宠爱也好，以及朋友的陪伴也好，本来都应该是她的。但她现在却只能在无人问津的角落里，依靠网络连接的手机去小心翼翼地偷瞄几眼原属于自己的生日会。​", ""}
	txt15	= []string{" “你看V魔怔了，真恶心。”　　看着同学发给你的消息，你陷入了沉思。　　仔细想一想，你觉得自己确实魔怔了，即使被鄙夷想要向熟人安利向晚小姐。　　思考再三过后，你决定回到现实，放弃入脑。　　今晚有向晚小姐的直播，你狠下心，没有点进去，而是倒头就睡。　　第二天，你起得很早。　　因为自从你开始看向晚小姐后，就养成了早睡早起的习惯，再也没有赖床过。　　你离开了狭小的出租屋，来到了公司，投入了工作中。　　工作很累，你感觉有些疲惫。　　你想起了嘉然小姐出道视频的不堪评论，以及她的笑容，烦闷减轻了不少。　　最近组长夸奖你工作很努力，别人不知道为何一向懒惰的你，在几个星期前开始一反常态地勤奋。　　只有你自己知道理由。　　撑过了加班的时间，你回到了出租屋，打开了外卖软件，却发现会员已经断了，你舍不得那些钱，所以决定自己买菜做饭。　　实际上，你看向晚小姐后就开始第一次尝试着做饭，没有点过外卖了。　　在超市你看到自己一向很喜欢的薯片在打折，愣了一下，没有买。　　因为你在看向晚小姐后，就再也没有暴饮暴食过，甚至连零食都戒掉。　　回家把饭做好，你安静地吃完了。　　看着电脑，你发现游戏已经很久没更新了。　　因为你在看向晚小姐后，也把一直沉迷到通宵的游戏给戒掉了。　　最终，你还是打开昨天晚上的录播，看到了向晚小姐热情地打招呼。　　“顶晚人们，晚上好呀~！”向晚小姐元气地打招呼。　　你发自内心地笑了起来。　　“晚上好！”你说。　　你忽然意识到一件事，她从来没有在你现实里出现过，却已经将你糟糕的生活改变。　　……　　看完录播后，你打开了贴吧，看到了熟悉的顶碗人在用表情包引流，有不明真相的路人在骂饭圈贵物。　　你无视了那些辱骂声，熟练地在下面RP，回复道：　　“这是我爹？好可爱呀，她的名字是什么？” ", ""}
	txt16	= []string{"你饿了，她有一块饼。再去买一块和你一起吃，这是贝拉；和你一人一半吃个半饱，这是珈乐；全部给你不舍得你挨饿，然后告诉你自己吃过了，这是嘉然；把饼偷偷扔掉，跟你一块饿，这是乃琳；把饼藏起来想给你一个惊喜，但是掏出来时发现碎成渣吃不了了，于是化身小哭包的，这是我的铸币晚晚。​", ""}
	txt17	= []string{"“你对嘉然小姐的爱有多重？”“大约300克”“300克？你是想说人类的心脏大约是300克吗？”“不，鼠鼠的平均体重大约是300克，因此我是全身心地爱着嘉然小姐。”[给心心][给心心]​", ""}
	txt18	= []string{"你身体里的每一个原子都来自一颗爆炸了的恒星，你左手的原子与右手的原子也许来自不同的恒星。这实在是我所知道的物理学中最富诗意的东西： 你的一切都是星尘......这其中经历了多少，惊险、巧合，三千世界的交错、重叠，才能让我们遇见彼此。我们都是星尘啊。但是，我这样的人，是星尘里，黯淡无光的那种。即使是星，也应该是最不闪亮的那一颗。而然然，你不一样。你是天上的星辰，闪耀而夺目。即使宇宙里，有着不计其数的、炫目的一等星，也远不如你的光辉。轻声哼唱，众星敛了光芒为你聆听。哭起鼻子，傍晚的霞也因你黯然。跳一支舞，似月兔回旋于空中。说起情话，登时小雨疏疏、浸润心窝。你是万千童话里，被守护的女主角。我喜欢你，像喜欢天上的星辰那样喜欢你。​", ""}
	txt19	= []string{"嘉然回来了，在枝江的大桥。她回来的那天是六月，天上却飘着雪。有人告诉我，她可能有还未了却的心愿或是什么怨气。我知道， 可能我做不到。第一次见到她，是在枝江。她带着甜甜的笑看着我，看着我有些发怵。有那么一瞬我似乎看到了阳光。我慌乱的移开视线，她径直朝我走来，微微弯着腰，面带戏虐的望着我。久居下水道的老鼠，第一次看见光，是睁不开眼的。 我慌乱的想赶紧逃走。“小老鼠，你好呀。”那是她对我说的第一句话。[给心心][给心心][给心心]​", ""}
	txt20	= []string{"听了好多遍，真的很震撼，作者属实是用心了，今天就不发病了，第一次写长评，整点读后感。直播里，“背靠着早已雪停的窗棂，你说着未曾到达的山顶”，那是与你失之交臂的舞台梦 。“我装着第三人称的淡定，用风轻云淡的态度掩饰那段经历”，我已经记不起她的姓名，只记得起舞翩翩的侧影”，过去那个满腔热情，在舞台上挥洒汗水的你已经淡出记忆，只依稀留下曾经翩翩起舞的身影。命不是总有天道酬情的注定，那么公平。台下十年苦练，只为台上耀眼的一刹，可当腰伤的意外来袭，梦却碎了。躺在病床上，难道梦想就要在此终结了吗？这也罢，也罢。那个曾经的你，也经历过数个春秋，经历汗水与欢乐，但她却走过了如红楼梦的悲情，从盛极一时万家灯火走到遁入空门的万物凋零。时光抚平了曾经的伤疤，心中的芥蒂也如沙石被流过的光阴慢慢冲刷消散，女孩又重新拾起了向往舞台的梦想。终于，她在asoul找到了新的舞台，遇到了支持她的一个魂们。“我叫贝拉”此刻，你向全世界宣布你的重生！贝极星们透过屏幕，看你在镜头前如此快乐自信，这让我也重拾梦想去直面未知的未来。为你的坚强而动容，在直播间为你打call呐喊，字字句句都发自真心。“我们闪烁在夜空，想照亮你夜晚的梦，没月亮时你可向北方转动，我们在天空”，追梦的路上，贝极星们陪着你分享每一次快乐，和你经历每一个坎坷，你若是感到失望，那就抬头看向天空吧，我们一直都在。女孩站在那个梦想中善良闪亮的舞台上，再回望过去，时光流过宛若江河，她住在江头，而那个过去的她住在江尾，回味着过去，竟有些恍如隔世的朦胧。但往者不可谏，来者犹可追，愿你携手曾经那个满怀梦想的自己，全力以赴，去追寻自己的Asoul梦吧，贝极星们会一直陪着你。​", ""}
	txt21	= []string{"不推嘉然小姐十年了。她的名气和出场费都一涨再涨，我原地踏步的工资买不上专辑也打不起榜。终于年前被公司安排下岗，找工作时我才在路边广告发现初代工具人已经当上了厂长。时间太久，一切都变了。到处投简历的时候我想起了一名人上人的预言：“这些人只配在下水道里度过相对比较失败的人生。”像是一条跳过龙门的锦鲤，金鳞被羽耀武扬威地站在门沿上，对其他还在跳的鲤鱼说：“你不行！”我当时很想反驳，可他说中了。我知道我确实不行。我之所以跳了跳，只是为了看下自己能跳成什么样罢了。其实每条鲤鱼的龙门都不是一样高的。我见过龙门在水下的鲤鱼。看起来是鱼，其实生而为龙。也有的生而为鱼肉。也见过好运的鲤鱼，门被各种大手摁到河里了。我也期待过好运，只是没来而已。说起来这就是人性吧。我不讨厌天道酬勤，但是讨厌别人的好运——只是因为我没有好运罢了。我也有亲人和宠物会生病；我眼神也挺纯真啊。讨厌嘉然小姐十年了。讨厌的更是越来越深的无力感。身在泥潭的人是没力气冲锋的吧。三流的人生只会让上等人不屑一顾吧。我坐井观天，天穹星海依然耀眼。可我爬不出井底。那我就不再看星星了。世界那么大，但没我的份。忘记嘉然小姐十年了。可路上看见街边的大荧幕在放A-soul的新年节目，我还是楞在那里了。我没有近视，但总觉得眼睛影影绰绰，雾气来自多年以前。这个广告位非常贵。真的再也不是小v了啊。抖友还在惦记他们的鸭子。晚晚仍然只有蓬蓬裙，100首歌竟然还欠着，被粉头小团体以4%年化复利计在小本子了。想起她首播时玩2077下饭下得轰轰烈烈，我一边发“粉丝牌改成晚饭人吧”“和嘉然珈乐凑加碗饭”“和乃琳凑来碗饭”一边忍住刷“和贝拉组拉碗饭”的冲动。solo依然拉跨，参团照旧神C。贝拉总是六边形战士，乃琳养了成吨的gachi，珈乐还是那个硬壳软妹。嘉然小姐依然卖萌摁混。什么都没变，是我没跟上她们。城里烟火幢幢，灯光下的人热情相拥，阴影里的人压下悸动。最亮的地方嘉然小姐浅笑起舞，光影从她袖间散落，像是雨天花伞轻旋，摇曳间洒下泪色的流珠。忽然眼睛有点模糊。我小声说：“新年好啊，嘉然小姐。”不爱嘉然小姐十年了。十年里，爱过的每个人都像她。", ""}
	txt22	= []string{"喜欢然然，不加掩饰总是微笑的然然，不是妖治妩媚的华丽舞姬，不是残忍无情的冷酷公主，只有着单纯烂漫的花样笑容，不带忧伤，却如同逝零而来的天际之风，飘卷了我心上的忧郁。棕色的长发装饰着美丽的蝴蝶结，宝石般的双眸中隐隐流露不存在的神伤，就是这样子的然然，不加掩饰总是微笑的然然，拥有烂漫笑容的然然，拥有亘古不变的美丽。说不出的感觉，却知道因为然然的笑容，化开了初春冰雪的痕迹。想要伸出手抚摸然然的发丝，可是却只能触到冰冷的硬屏，只有然然清澈的眼神，依旧是我所熟悉。然然永远清澈的瞳仁然然烂漫单纯的笑容然然亘古不变的美丽然然是我唯一的最爱，然然是我最爱的唯一。然然的喜好、快乐、难过、伤心和无助，我都了如指掌，喜欢草莓蛋糕和黄色，个性温柔又善良，可爱体贴的女孩子，有着说不出的优点，我想然然一定是上天赐予我最好的礼物。"}
)


func init() {
	// 随机发送一篇上面的小作文
	zero.OnFullMatch("/小作文").
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(utils.RandText(txt1,txt2,txt3,txt4,txt5,txt6,txt7,txt8,txt9,txt11,txt12,txt13,txt14,txt15,txt16,txt17,txt18,txt19,txt20,txt21,txt22))
	})

	// 逆天
	zero.OnFullMatch("发大病", zero.OnlyToMe).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("贝拉抽我\U0001F975 嘉然骑在我背上 \U0001F975晚晚踩我\U0001F975 乃琳坐在王座是用看垃圾的眼神看我 \U0001F975\U0001F975珈乐踢我\U0001F975\U0001F975，把我眼睛蒙住然后五只脚一起踩我\U0001F975还让我猜脚是谁的，猜错了给我劈眼一铁棍\U0001F975​")
		})

	// 这其实是一个炸群功能
	zero.OnFullMatch("随机函数测试", zero.OnlyToMe, zero.AdminPermission).
		Handle(func(ctx *zero.Ctx) {
			for i:=1; i<1000; i++ {
				ctx.Send(utils.Suiji())
				time.Sleep(300 * time.Millisecond)
			}
	})
}
