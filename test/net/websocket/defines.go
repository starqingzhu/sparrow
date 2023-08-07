package websocket

const (
	PACKET_GWC_NONE                           = iota + 61000 // 0£¬¿Õ
	PACKET_CGW_VERIFY_PAK                                    //Client Connect to GateServer
	PACKET_GWC_VERIFY_PAK                                    //Client Connect to GateServer
	PACKET_CGW_HEART_PAK                                     //Client Connect to GateServer
	PACKET_CGW_CONNECT_GAMESERVER_PAK                        //Client Connect to GateServer
	PACKET_GWC_CONNECT_GAMESERVER_RET_PAK                    //Client Connect to GateServer
	PACKET_CGW_REQUEST_GAMESERVERINFO_PAK                    //Client Connect to GateServer
	PACKET_GWC_REQUEST_GAMESERVERINFO_RET_PAK                //Client Connect to GateServer
	PACKET_CGW_RECONNECT_GAMESERVER_PAK                      //client reconnect
	PACKET_GWC_HEART_PAK                                     //Client Connect to GateServer
	PACKET_CGW_CLIENT_HEART_PAK                              //Client Connect to GateServer
	PACKET_GWC_CLIENT_HEART_PAK                              //Client Connect to GateServer
	PACKET_CGW_CLIENT_LOGIN_PAK                              //Client Send Login to GateServer
	PACKET_CGW_QUERY_GATE_INFO_PAK                           //Client Query Gate Info
	PACKET_GWC_QUERY_GATE_INFO_RET_PAK                       //Gate Return Info To Client

	PACKET_GWC_MAX //消息类型的最大值
	PACKET_GWC_NUM = PACKET_GWC_MAX - PACKET_GWC_NONE - 1
)

const (
	PACKET_CP_NONE                                         = iota + 300 // 0，空
	PACKET_CP_LOGIN_PAK                                                 //client ask login
	PACKET_PC_LOGIN_RET_PAK                                             //client login result
	PACKET_PC_ENTER_WORLD_PAK                                           //client login result
	PACKET_CP_CLIENT_CONFIG_PAK                                         //save client config
	PACKET_CP_RECONNECT_PAK                                             //
	PACKET_PC_RECONNECT_RET_PAK                                         //
	PACKET_CP_GM_COMMAND_PAK                                            //
	PACKET_PC_GM_COMMAND_PAK                                            //
	PACKET_CP_CREATE_PLAYER_PAK                                         //
	PACKET_PC_CREATE_PLAYER_RET_PAK                                     //
	PACKET_CP_ACCOUNT_BIND_PAK                                          //
	PACKET_PC_ACCOUNT_BIND_PAK                                          //
	PACKET_CP_CLIENT_READY_PAK                                          //
	PACKET_CP_CONNECTED_HEARTBEAT_PAK                                   //client connected heartbeat
	PACKET_PC_KICK_OUT_PAK                                              //
	PACKET_PC_SYN_PLAYER_INFO_PAK                                       //
	PACKET_PC_NOTICE_PAK                                                //PC_NOTICE
	PACKET_CP_CHANGE_EQUIPMENT_PAK                                      //
	PACKET_PC_CHANGE_EQUIPMENT_PAK                                      //
	PACKET_PC_ITEM_LIST_PAK                                             //PC_ITEM_LIST
	PACKET_PC_ITEM_CHANGE_PAK                                           //PC_ITEM_CHANGE
	PACKET_CP_CARD_COMPOSE_PAK                                          //CP_CARD_COMPOSE
	PACKET_PC_CARD_COMPOSE_PAK                                          //PC_CARD_COMPOSE
	PACKET_CP_CARD_EQUIP_PAK                                            //CP_CARD_EQUIP
	PACKET_PC_CARD_EQUIP_PAK                                            //PC_CARD_EQUIP
	PACKET_CP_CARD_TAKEOFF_PAK                                          //CP_CARD_TAKEOFF
	PACKET_PC_CARD_TAKEOFF_PAK                                          //PC_CARD_TAKEOFF
	PACKET_CP_CARD_UPGRADE_PAK                                          //CP_CARD_UPGRADE
	PACKET_PC_CARD_UPGRADE_PAK                                          //PC_CARD_UPGRADE
	PACKET_CP_CARD_DECOMPOSE_PAK                                        //CP_CARD_DECOMPOSE
	PACKET_PC_CARD_DECOMPOSE_PAK                                        //PC_CARD_DECOMPOSE
	PACKET_CP_CARD_FORGE_PAK                                            //CP_CARD_FORGE
	PACKET_PC_CARD_FORGE_PAK                                            //PC_CARD_FORGE
	PACKET_CP_CARD_FORGE_LOCK_PAK                                       //CP_CARD_FORGE_LOCK
	PACKET_PC_CARD_FORGE_LOCK_PAK                                       //PC_CARD_FORGE_LOCK
	PACKET_CP_FASHION_CHANGE_PAK                                        //
	PACKET_PC_FASHION_CHANGE_PAK                                        //
	PACKET_CP_CARD_COMPOSECHIP_PAK                                      //CP_CARD_COMPOSECHIP
	PACKET_PC_CARD_COMPOSECHIP_PAK                                      //PC_CARD_COMPOSECHIP
	PACKET_PC_ITEM_NOTICE_LIST_PAK                                      //PC_ITEM_NOTICE_LIST
	PACKET_CP_ITEM_SALE_PAK                                             //CP_ITEM_SALE
	PACKET_PC_ITEM_SALE_PAK                                             //PC_ITEM_SALE
	PACKET_CP_ITEM_USE_PAK                                              //CP_ITEM_USE
	PACKET_PC_ITEM_USE_PAK                                              //PC_ITEM_USE
	PACKET_PC_STATUS_LIST_PAK                                           //PC_STATUS_LIST
	PACKET_CP_CHANGE_COLOR_PAK                                          //
	PACKET_PC_CHANGE_COLOR_PAK                                          //
	PACKET_CP_CONFIRM_CHANGE_COLOR_PAK                                  //
	PACKET_PC_CONFIRM_CHANGE_COLOR_PAK                                  //
	PACKET_PC_EMAIL_LIST_PAK                                            //PC_EMAIL_LIST
	PACKET_PC_EMAIL_INFO_PAK                                            //PC_EMAIL_INFO
	PACKET_CP_EMAIL_READ_PAK                                            //CP_EMAIL_READ
	PACKET_PC_EMAIL_READ_PAK                                            //PC_EMAIL_READ
	PACKET_CP_EMAIL_AWARD_PAK                                           //CP_EMAIL_AWARD
	PACKET_PC_EMAIL_AWARD_PAK                                           //PC_EMAIL_AWARD
	PACKET_CP_EMAIL_DEL_PAK                                             //CP_EMAIL_DEL
	PACKET_PC_EMAIL_DEL_PAK                                             //PC_EMAIL_DEL
	PACKET_CP_ACTIVITY_INFO_PAK                                         //
	PACKET_PC_ACTIVITY_INFO_PAK                                         //
	PACKET_CP_RECEIVE_AWARD_PAK                                         //
	PACKET_PC_RECEIVE_AWARD_PAK                                         //
	PACKET_CP_BACK_TO_LOBBY_PAK                                         //
	PACKET_CP_GOODS_INFO_PAK                                            //CP_GOODS_INFO
	PACKET_PC_GOODS_INFO_PAK                                            //PC_GOODS_INFO
	PACKET_CP_GOODS_BUY_PAK                                             //CP_GOODS_BUY
	PACKET_PC_GOODS_BUY_PAK                                             //PC_GOODS_BUY
	PACKET_CP_SHOP_REFRESH_PAK                                          //shop refresh
	PACKET_PC_SHOP_REFRESH_RET_PAK                                      //shop refresh ret
	PACKET_PC_TASK_LIST_PAK                                             //PC_TASK_LIST
	PACKET_PC_TASK_STARAWARD_LIST_PAK                                   //PC_TASK_STARAWARD_LIST
	PACKET_CP_TASK_STARAWARD_AWARD_PAK                                  //CP_TASK_STARAWARD_AWARD
	PACKET_PC_TASK_STARAWARD_AWARD_PAK                                  //PC_TASK_STARAWARD_AWARD
	PACKET_CP_TASK_STARAWARD_LVLUP_PAK                                  //CP_TASK_STARAWARD_LVLUP
	PACKET_PC_TASK_STARAWARD_LVLUP_PAK                                  //PC_TASK_STARAWARD_LVLUP
	PACKET_CP_TASK_STARAWARD_BUYSTAR_PAK                                //CP_TASK_STARAWARD_BUYSTAR
	PACKET_PC_TASK_STARAWARD_BUYSTAR_PAK                                //PC_TASK_STARAWARD_BUYSTAR
	PACKET_CP_TASK_DAILY_REFRESH_PAK                                    //CP_TASK_DAILY_REFRESH
	PACKET_PC_TASK_DAILY_REFRESH_PAK                                    //PC_TASK_DAILY_REFRESH
	PACKET_CP_SEVENLOGIN_AWARD_PAK                                      //CP_SEVENLOGIN_AWARD
	PACKET_PC_SEVENLOGIN_AWARD_PAK                                      //PC_SEVENLOGIN_AWARD
	PACKET_PC_SEVENLOGIN_AWARD_INFO_PAK                                 //PC_SEVENLOGIN_AWARD_INFO
	PACKET_CP_AFTER_SIGNIN_PAK                                          //Supplementary signature
	PACKET_PC_AFTER_SIGNIN_RET_PAK                                      //back result for Supplementary signature
	PACKET_CP_ENROLL_MATCH_PAK                                          //enroll match
	PACKET_PC_ENROLL_MATCH_RET_PAK                                      //enroll match result
	PACKET_PC_ENTER_MATCH_ROOM_PAK                                      // match room
	PACKET_PC_WEEKCARDITEM_INFO_PAK                                     //PC_WEEKCARDITEM_INFO
	PACKET_PC_NOTIFYEQUIPMENTUNLOCKLIST_PAK                             //
	PACKET_CP_REQUESTRECHARGELIST_PAK                                   //
	PACKET_PC_REQUESTRECHARGELISTRET_PAK                                //
	PACKET_CP_LUCKYWHEEL_INFO_PAK                                       //CP_LUCKYWHEEL_LIST
	PACKET_PC_LUCKYWHEEL_INFO_RESULT_PAK                                //PC_LUCKYWHEEL_INFO_RESULT
	PACKET_CP_PICK_LUCKYWHEEL_PAK                                       //CP_PICK_LUCKYWHEEL
	PACKET_PC_PICK_LUCKYWHEEL_RESULT_PAK                                //PC_PICK_LUCKYWHEEL_RESULT
	PACKET_CP_RECHARGE_FIRST_REWARD_PAK                                 //CP_RECHARGE_FIRST_REWARD
	PACKET_PC_RECHARGE_FIRST_REWARD_PAK                                 //PC_RECHARGE_FIRST_REWARD
	PACKET_PC_SYNC_COMMONDATA_PAK                                       //sys com data
	PACKET_PC_SYNC_COMMONFLAG_PAK                                       //sys com flag
	PACKET_PC_SYNC_COMMONFLAGBIT_PAK                                    //sys com flag
	PACKET_CP_BATCH_SALE_PAK                                            //
	PACKET_PC_BAGTIP_PAK                                                //
	PACKET_CP_EXCHANGE_JJGOLD_PAK                                       //
	PACKET_PC_EXCHANGE_JJGOLD_RESULT_PAK                                //
	PACKET_CP_QUERYJJGOLD_PAK                                           //
	PACKET_PC_QUERYJJGOLD_PAK                                           //
	PACKET_CP_REQUESTMONSTERSCORERECORD_PAK                             //CP_RequestMonsterScore
	PACKET_PC_REQUESTMONSTERSCORERECORD_PAK                             //PC_RequestMonsterScore
	PACKET_CP_REQUESTMONSTERSCOREINTERVAL_PAK                           //CP_RequestMonsterScoreInterval
	PACKET_PC_REQUESTMONSTERSCOREINTERVAL_PAK                           //PC_RequestMonsterScoreInterval
	PACKET_CP_USETREASURECHESTHAMMER_PAK                                //
	PACKET_PC_USETREASURECHESTHAMMER_PAK                                //
	PACKET_CP_ENHANCEEQUIP_PAK                                          //
	PACKET_PC_ENHANCEEQUIP_PAK                                          //
	PACKET_PC_EQUIPSUITINFO_PAK                                         //
	PACKET_CP_RECHARGEVERIFY_PAK                                        //
	PACKET_PC_RECHARGEVERIFY_PAK                                        //PC
	PACKET_CP_REQUESTRECHARGEORDERID_PAK                                //
	PACKET_PC_REQUESTRECHARGEORDERIDRET_PAK                             //
	PACKET_PC_RECHARGESUC_PAK                                           //
	PACKET_CP_SUCRECHARGEORDERID_PAK                                    //
	PACKET_CP_EXCHANGEHAMMER_PAK                                        //
	PACKET_PC_EXCHANGEHAMMER_PAK                                        //
	PACKET_CP_ENHANCEATTR_PAK                                           //enhance attr
	PACKET_PC_ENHANCEATTR_RET_PAK                                       //enhance attr
	PACKET_PC_JACKPOTHAVELOTTRYTASKREWARD_PAK                           // syn Jackpot have lottrynum task reward
	PACKET_CP_JACKPOTLOTTRYINFO_PAK                                     //Jackpot lottry num and task state
	PACKET_PC_JACKPOTLOTTRYINFO_RET_PAK                                 //Jackpot lottry num and task state ret
	PACKET_CP_JACKPOTLOTTRYNUMREWARD_PAK                                //get jackpot lottrynum task reward
	PACKET_PC_JACKPOTLOTTRYNUMREWARD_RET_PAK                            //get jackpot lottrynum task reward ret
	PACKET_CP_JACKPOTLOTTRY_PAK                                         //Jackpot lottry
	PACKET_PC_JACKPOTLOTTRY_RET_PAK                                     //Jackpot lottry ret
	PACKET_CP_JACKPOTLOTTRY_POOLITEMNUM_PAK                             //Jackpot lottry pool item num
	PACKET_PC_JACKPOTLOTTRY_POOLITEMNUM_RET_PAK                         //Jackpot lottry pool item num ret
	PACKET_CP_JACKPOTLOTTRY_RECORD_PAK                                  //Jackpot lottry record
	PACKET_PC_JACKPOTLOTTRY_RECORD_RET_PAK                              //Jackpot lottry record ret
	PACKET_CP_TICKET_LOTTERY_PAK                                        //ticket lottery
	PACKET_PC_TICKET_LOTTERY_RET_PAK                                    //ticket lottery ret
	PACKET_CP_TICKETLOTTERY_RECORD_PAK                                  //get ticket lottery record
	PACKET_PC_TICKETLOTTERY_RECORD_RET_PAK                              //get ticket lottery record ret
	PACKET_CP_MONOPOLYINFO_PAK                                          //Monopoly Activity Info
	PACKET_PC_MONOPOLYINFO_RET_PAK                                      //Monopoly Activity Info ret
	PACKET_CP_MONOPOLYROLLDICE_PAK                                      //get Monopoly Roll Dice
	PACKET_PC_MONOPOLYROLLDICE_RET_PAK                                  //get Monopoly Roll Dice ret
	PACKET_CP_MONOPOLYSTOREYTASKREWARD_PAK                              //get Monopoly  Storey task reward
	PACKET_PC_MONOPOLYSTOREYTASKREWARD_RET_PAK                          //get Monopoly  Storey task reward ret
	PACKET_PC_MONOPOLYHAVESTOREYTASK_PAK                                // syn Monopoly have task reward
	PACKET_PC_TREASURELAND_ICONNOTICE_PAK                               //TreasureLand Activity IconNotice
	PACKET_CP_TREASUREINFO_PAK                                          //TreasureLand Activity Info
	PACKET_PC_TREASUREINFO_RET_PAK                                      //TreasureLand Activity Info ret
	PACKET_CP_TREASURELAND_REWARD_PAK                                   //TreasureLand Get LandType RewardList
	PACKET_PC_TREASURELAND_REWARD_PAK                                   //TreasureLand Get LandType RewardList ret
	PACKET_CP_TREASURELAND_MOVE_PAK                                     //TreasureLand Move
	PACKET_PC_TREASURELAND_MOVE_RET_PAK                                 //TreasureLand Get LandType RewardList ret
	PACKET_CP_TREASURELAND_END_PAK                                      //TreasureLand End
	PACKET_PC_TREASURELAND_END_RET_PAK                                  //TreasureLand Get LandType RewardList ret
	PACKET_CP_TREASURE_CHIPEXCHANGE_PAK                                 //TreasureLand Chip exchange
	PACKET_PC_TREASURE_CHIPEXCHANGE_RET_PAK                             //TreasureLand Chip exchange ret
	PACKET_CP_COMPOUND_PAK                                              //Compound
	PACKET_PC_COMPOUNDRESULT_PAK                                        //Compound
	PACKET_CP_DAILYTASK_INFO_PAK                                        //Get Daily task info
	PACKET_PC_DAILYTASK_INFO_RET_PAK                                    //Get Daily task info
	PACKET_PC_SYN_DAILYTASK_PROGRESS_PAK                                //syn Daily task Progress
	PACKET_CP_DAILYTASK_ACCEPT_PAK                                      //accept Daily task
	PACKET_PC_DAILYTASK_ACCEPT_RET_PAK                                  //accept Daily task ret
	PACKET_CP_DAILYTASK_REFRESH_PAK                                     //refresh Daily task
	PACKET_PC_DAILYTASK_REFRESH_RET_PAK                                 //refresh Daily task
	PACKET_CP_DAILYTASK_RECORD_PAK                                      //get Daily task  record
	PACKET_PC_DAILYTASK_RECORD_RET_PAK                                  //get Daily task  record ret
	PACKET_PC_COMMON_TIPS_PAK                                           //common tip
	PACKET_PC_SYN_MODULELOCK_DATA_PAK                                   //module lock syn info鍔熻兘瑙ｉ攣鏁版嵁鍚屾
	PACKET_PC_SCRAPCARD_SYNCINFO_PAK                                    //guaguaka - award pool info
	PACKET_PC_SCRAPCARD_OWNCARDCOUNT_PAK                                //guaguaka - player exchanged count
	PACKET_CP_SCRAPCARD_POOLEXCHANGECOUNTREQ_PAK                        //guaguaka - req pool exchange count
	PACKET_PC_SCRAPCARD_POOLEXCHANGECOUNT_PAK                           //guaguaka - update pool exchanged count
	PACKET_CP_SCRAPCARD_EXCHANGEREQ_PAK                                 //guaguaka - req exchange card
	PACKET_PC_SCRAPCARD_EXCHANGERET_PAK                                 //guaguaka - exchange card ret
	PACKET_CP_SCRAPCARD_AWARDINFOREQ_PAK                                //guaguaka - req award info
	PACKET_PC_SCRAPCARD_AWARDINFORET_PAK                                //guaguaka - return award info
	PACKET_CP_SCRAPCARD_AWARDLOGREQ_PAK                                 //guaguaka - req award log
	PACKET_PC_SCRAPCARD_AWARDLOGRET_PAK                                 //guaguaka - return award log
	PACKET_PC_SCRAPCARD_CARDLIST_PAK                                    //guaguaka - player card list
	PACKET_CP_SCRAPCARD_SCRAPREQ_PAK                                    //guaguaka - req scrap card
	PACKET_PC_SCRAPCARD_SCRAPRET_PAK                                    //guaguaka - scrap card ret
	PACKET_CP_SCRAPCARD_SCRAPCOMMONREQ_PAK                              //guaguaka - onekey scrap common card
	PACKET_PC_SCRAPCARD_SCRAPCOMMONRET_PAK                              //guaguaka - onekey scrap common card result
	PACKET_CP_SCRAPCARD_STARTSCRAP_PAK                                  //guaguaka - start scrap
	PACKET_CP_SCRAPCARD_STOPSCRAP_PAK                                   //guaguaka - stop scrap
	PACKET_PC_SCRAPCARD_SETTLEMENT_PAK                                  //guaguaka - settlement
	PACKET_PC_SCRAPCARD_HAVERECORDLOG_PAK                               //guaguaka - have record log
	PACKET_PC_SCRAPCARD_CLOSE_PAK                                       //guaguaka - close
	PACKET_PC_SCRAPCARD_OPEN_PAK                                        //guaguaka - open
	PACKET_CP_STORY_MISSION_INFO_PAK                                    //get mission info
	PACKET_PC_STORY_MISSION_INFO_PAK                                    // syn mission info
	PACKET_CP_GET_STORY_MISSION_REWARDS_PAK                             // get mission reward
	PACKET_PC_GET_STORY_MISSION_REWARDS_RET_PAK                         // get mission reward ret
	PACKET_CP_MISSIONCLIENTEVENT_PAK                                    //mission client event
	PACKET_PC_MISSIONCLIENTEVENTRET_PAK                                 // mission client event ret
	PACKET_CP_STARTBUILD_PAK                                            // homland build
	PACKET_PC_STARTBUILD_RET_PAK                                        // homland build
	PACKET_CP_ACTIVATEBUILD_PAK                                         // homland Activate build
	PACKET_PC_ACTIVATEBUILD_RET_PAK                                     // homland Activate build
	PACKET_PC_BUILDSTATUS_PAK                                           // syn building status
	PACKET_CP_BUILDCOMPOSE_PAK                                          // build Compose
	PACKET_PC_BUILDCOMPOSE_RET_PAK                                      // build Compose ret
	PACKET_CP_BUILDMATERIALREWARD_PAK                                   // build material
	PACKET_PC_BUILDMATERIALREWARD_RET_PAK                               // build material ret
	PACKET_CP_PET_HATCH_PAK                                             // hatch pet
	PACKET_CP_PET_EVOLVE_PAK                                            // evovle pet
	PACKET_CP_PET_SUMMON_PAK                                            // summon pet
	PACKET_CP_PET_REST_PAK                                              // rest pet
	PACKET_CP_PET_FREE_PAK                                              // free pet
	PACKET_PC_PET_HATCHRET_PAK                                          // hatch pet result
	PACKET_PC_PET_SUMMON_RET_PAK                                        // summon pet result
	PACKET_PC_PET_REST_RET_PAK                                          // rest pet result
	PACKET_PC_PET_EVOLVE_RET_PAK                                        // evovle pet result
	PACKET_PC_PET_FREE_RET_PAK                                          // free pet result
	PACKET_PC_PET_LIST_PAK                                              // pet list
	PACKET_PC_PET_CURPET_PAK                                            // cur pet
	PACKET_PC_PET_HATCHING_PAK                                          // cur hatching pet
	PACKET_PC_PET_HATCHOVER_PAK                                         // hatch over
	PACKET_PC_CLIENT_CONFIG_PAK                                         // PC_CLIENT_CONFIG
	PACKET_CP_CHANGE_TITLE_PAK                                          //CP_ChangeTitle
	PACKET_PC_CHANGE_TITLE_RET_PAK                                      //PC_ChangeTitle
	PACKET_PC_TITLE_LIST_PAK                                            //PC_Title_LIST
	PACKET_PC_SYN_DEL_TITLE_PAK                                         //PC_SYN_DEL_TITLE
	PACKET_PC_SYN_ADD_TITLE_PAK                                         //PC_SYN_ADD_TITLE
	PACKET_CP_USE_CONSUM_ITEM_PAK                                       //CP_USE_CONSUM_ITEM
	PACKET_PC_USE_CONSUM_ITEM_PAK                                       //PC_USE_CONSUM_ITEM
	PACKET_PC_SYN_LIMITTIMES_PAK                                        //PC_Syn_LimitTimes
	PACKET_PC_FUNCTIONPETLIST_PAK                                       //
	PACKET_PC_FUNCTIONPET_NEW_PAK                                       //
	PACKET_CP_FUNCTIONPETCONFIG_PAK                                     //
	PACKET_PC_FUNCTIONPETCONFIG_PAK                                     //
	PACKET_CP_SUMMONFUNCTIONPET_PAK                                     //
	PACKET_PC_SUMMONFUNCTIONPET_RET_PAK                                 //
	PACKET_CP_RESTFUNCTIONPET_PAK                                       //
	PACKET_PC_RESTFUNCTIONPET_RET_PAK                                   //
	PACKET_PC_FUNCTIONPET_EXPIRE_PAK                                    //
	PACKET_PC_SYN_ATTR_PAK                                              //Server Syn Attr
	PACKET_PC_SYN_SYSTEMRECRODSTATE_PAK                                 //
	PACKET_CP_GET_SYSTEMRECROD_PAK                                      //
	PACKET_PC_GET_SYSTEMRECROD_PAK                                      //
	PACKET_CP_SIGNACTIVITY_INFO_PAK                                     //Get sign activity info
	PACKET_PC_SIGNACTIVITY_INFO_RET_PAK                                 //Get sign activity info
	PACKET_PC_SYN_SIGNACTIVITYTASK_PROGRESS_PAK                         //syn sign activity task Progress
	PACKET_CP_TRANSFERCAREER_PAK                                        //Transfer Career
	PACKET_PC_TRANSFERCAREER_RET_PAK                                    //Transfer Career
	PACKET_CP_RECHARGECONSUME_PAK                                       //
	PACKET_PC_RECHARGEREBATE_INFO_PAK                                   //
	PACKET_PC_RECHARGEREBATE_SELFEXCHANGEDCOUNT_PAK                     //
	PACKET_PC_RECHARGEREBATE_SELFAWARDLIST_PAK                          //
	PACKET_PC_RECHARGEREBATE_POOLEXCHANGEDCOUNT_PAK                     //
	PACKET_PC_RECHARGEREBATE_TOPAWARDLIST_PAK                           //
	PACKET_PC_RECHARGEREBATE_HAVEAWARD_PAK                              //
	PACKET_CP_RECHARGEREBATE_REQSELFAWARDLOG_PAK                        //
	PACKET_PC_RECHARGEREBATE_SELFAWARDLOGLIST_PAK                       //
	PACKET_CP_RECHARGEREBATE_REQEXCHANGE_PAK                            //
	PACKET_PC_RECHARGEREBATE_EXCHANGERESULT_PAK                         //
	PACKET_CP_RECHARGEREBATE_REQPOOLTOTALEXCHANGECOUNT_PAK              //
	PACKET_CP_RECHARGEREBATE_REQGETAWARD_PAK                            //
	PACKET_PC_RECHARGEREBATE_GETAWARDRESULT_PAK                         //
	PACKET_PC_RECHARGEREBATE_OPEN_PAK                                   //
	PACKET_PC_RECHARGEREBATE_CLOSE_PAK                                  //
	PACKET_CP_ONEKEYCHANGEFASHION_PAK                                   //
	PACKET_PC_ONEKEYCHANGEFASHIONRET_PAK                                //
	PACKET_CP_ONEKEYUPGRADEFASHION_PAK                                  //
	PACKET_PC_ONEKEYUPGRADEFASHIONRET_PAK                               //
	PACKET_CP_ONEKEYFUSION_PAK                                          //
	PACKET_PC_ONEKEYFUSIONRET_PAK                                       //
	PACKET_CP_SINGLEEQUIPMENTUPGRADE_PAK                                //
	PACKET_PC_SINGLEEQUIPMENTUPGRADERET_PAK                             //
	PACKET_CP_REQDRAW_PAK                                               //
	PACKET_CP_REQBATCHDRAW_PAK                                          //
	PACKET_PC_DRAWRESULT_PAK                                            //

	PACKET_CP_MAX //消息类型的最大值
)

const (
	PACKET_CG_NONE                              = iota + 0 // 0，空
	PACKET_CG_LOGIN_PAK                                    //client ask login
	PACKET_GC_LOGIN_RET_PAK                                //client login result
	PACKET_CG_RECONNECT_PAK                                //client reconnect
	PACKET_GC_RECONNECT_RET_PAK                            //client reconnect ret
	PACKET_CG_CONNECTED_HEARTBEAT_PAK                      //client connected heartbeat
	PACKET_GC_CONNECTED_HEARTBEAT_PAK                      //server connected heartbeat
	PACKET_GC_SYN_ATTR_PAK                                 //Server Syn Attr
	PACKET_CG_MOVE_PAK                                     //Player Move
	PACKET_GC_MOVE_PAK                                     //Notify Character Move
	PACKET_GC_STOP_PAK                                     //Notify Character Stop
	PACKET_GC_CHECK_POS_PAK                                //check pos by client
	PACKET_CG_GMCOMMAND_PAK                                //send gm command
	PACKET_GC_ENTER_SCENE_PAK                              //Enter Scene
	PACKET_CG_ENTER_SCENE_OK_PAK                           //Client Enter Scene OK
	PACKET_CG_LEAVE_GAME_PAK                               //
	PACKET_GC_LEAVE_GAME_PAK                               //
	PACKET_GC_GAME_FINISH_PAK                              //
	PACKET_GC_CREATE_PLAYER_PAK                            //Create Player
	PACKET_GC_DELETE_OBJ_PAK                               //Delete Player
	PACKET_CG_USE_SKILL_PAK                                //tel to server use skill
	PACKET_GC_USE_SKILL_PAK                                //broad cast use skill
	PACKET_GC_SKILL_BREAK_PAK                              //
	PACKET_CG_SKILL_FINISH_PAK                             //
	PACKET_GC_SKILL_FINISH_PAK                             //
	PACKET_GC_NOTIFY_CD_PAK                                //
	PACKET_GC_NOTIFY_SKILL_DIS_PAK                         //
	PACKET_GC_DAMAGE_INFO_PAK                              //Send DamageBoard Info to client
	PACKET_GC_BE_HIT_INFO_PAK                              //Send BeHit Info to client
	PACKET_GC_TRANSFORM_PAK                                //
	PACKET_GC_DIE_PAK                                      //game send die to client
	PACKET_CG_ASK_RELIVE_PAK                               //send Ask Relive
	PACKET_GC_SYN_SKILLINFO_PAK                            //Server Syn Skill Info
	PACKET_GC_RET_RELIVE_PAK                               //send Ret Relive
	PACKET_GC_ADD_BUFF_PAK                                 //GC_ADD_BUFF
	PACKET_GC_BUFF_COMMON_EFFECT_PAK                       //
	PACKET_GC_BUFF_EFFECT_PAK                              //
	PACKET_GC_DEL_BUFF_PAK                                 //GC_DEL_BUFF
	PACKET_GC_RELEASE_EFFECT_PAK                           //
	PACKET_GC_PASSIVE_SKILL_ADD_PAK                        //
	PACKET_GC_PASSIVE_SKILL_DEL_PAK                        //
	PACKET_GC_PASSIVE_SKILL_TRIGGER_PAK                    //
	PACKET_GC_COLLECT_STATUS_PAK                           //Collect status
	PACKET_CG_COLLECT_PAK                                  //CG_COLLECT
	PACKET_GC_COLLECT_PAK                                  //GC_COLLECT
	PACKET_GC_CREATE_COLLECTION_PAK                        //GC_CREATE_COLLECTION
	PACKET_GC_DEL_COLLECTION_PAK                           //GC_DEL_COLLECTION
	PACKET_GC_CREATE_OBJ_PAK                               //GC_CREATE_OBJ
	PACKET_CG_BUILD_CREATE_PAK                             //CG_BUILD_CREATE
	PACKET_GC_BUILD_CREATE_PAK                             //GC_BUILD_CREATE
	PACKET_GC_BUILD_SYN_STATUS_PAK                         //GC_BUILD_SYN_STATUS
	PACKET_CG_BUILD_CANCEL_PAK                             //CG_BUILD_CANCEL
	PACKET_GC_BUILD_CANCEL_PAK                             //GC_BUILD_CANCEL
	PACKET_GC_BUILD_SYN_LIST_PAK                           //GC_BUILD_SYN_LIST
	PACKET_CG_BUILD_UPGRADE_PAK                            //CG_BUILD_UPGRADE
	PACKET_GC_BUILD_UPGRADE_PAK                            //GC_BUILD_UPGRADE
	PACKET_CG_BUILD_REPAIR_PAK                             //CG_BUILD_REPAIR
	PACKET_GC_BUILD_REPAIR_PAK                             //GC_BUILD_REPAIR
	PACKET_CG_ITEM_MOVE_PAK                                //CG_ITEM_MOVE
	PACKET_GC_ITEM_MOVE_PAK                                //GC_ITEM_MOVE
	PACKET_CG_ITEM_TIDY_PAK                                //CG_ITEM_TIDY
	PACKET_GC_ITEM_TIDY_PAK                                //GC_ITEM_TIDY
	PACKET_CG_ITEM_DISCARD_PAK                             //CG_ITEM_DISCARD
	PACKET_GC_ITEM_DISCARD_PAK                             //GC_ITEM_DISCARD
	PACKET_CG_ITEM_DESTROY_PAK                             //CG_ITEM_DESTROY
	PACKET_GC_ITEM_DESTROY_PAK                             //GC_ITEM_DESTROY
	PACKET_CG_ITEM_USE_PAK                                 //CG_ITEM_USE
	PACKET_GC_ITEM_USE_PAK                                 //GC_ITEM_USE
	PACKET_CG_ITEM_REPAIR_PAK                              //CG_ITEM_REPAIR
	PACKET_GC_ITEM_REPAIR_PAK                              //GC_ITEM_REPAIR
	PACKET_CG_ITEM_PICKUP_PAK                              //CG_ITEM_PICKUP
	PACKET_GC_ITEM_PICKUP_PAK                              //GC_ITEM_PICKUP
	PACKET_GC_ITEM_SYN_STATUS_PAK                          //GC_ITEM_SYN_STATUS
	PACKET_GC_ITEM_LIST_PAK                                //GC_ITEM_LIST
	PACKET_GC_ITEM_CHANGE_PAK                              //GC_ITEM_CHANGE
	PACKET_CG_ITEM_CREATE_PAK                              //CG_ITEM_CREATE
	PACKET_GC_ITEM_CREATE_PAK                              //GC_ITEM_CREATE
	PACKET_CG_ITEM_SALE_PAK                                //CG_ITEM_SALE
	PACKET_GC_ITEM_SALE_PAK                                //GC_ITEM_SALE
	PACKET_CG_EQUIPMENT_LIST_PAK                           //CG_EQUIPMENT_LIST
	PACKET_GC_EQUIPMENT_LIST_PAK                           //GC_EQUIPMENT_LIST
	PACKET_CG_EQUIPMENT_DETAIL_PAK                         //CG_EQUIPMENT_DETAIL
	PACKET_GC_EQUIPMENT_DETAIL_PAK                         //GC_EQUIPMENT_DETAIL
	PACKET_CG_EQUIPMENT_TAKEOFF_PAK                        //CG_EQUIPMENT_TAKEOFF
	PACKET_GC_EQUIPMENT_TAKEOFF_PAK                        //GC_EQUIPMENT_TAKEOFF
	PACKET_GC_SYN_EQUIPMENT_APPEARANCE_PAK                 //GC_SYN_EQUIPMENT_APPEARANCE
	PACKET_GC_SCENE_NOTICE_PAK                             //GC_SCENE_NOTICE
	PACKET_CG_ITEM_STORAGE_PAK                             //CG_ITEM_STORAGE
	PACKET_GC_ITEM_STORAGE_PAK                             //GC_ITEM_STORAGE
	PACKET_CG_ITEM_FETCH_PAK                               //CG_ITEM_FETCH
	PACKET_GC_ITEM_FETCH_PAK                               //GC_ITEM_FETCH
	PACKET_CG_ENTER_BUILDING_PAK                           //CG_ENTER_BUILDING
	PACKET_GC_ENTER_BUILDING_PAK                           //GC_ENTER_BUILDING
	PACKET_GC_BOSS_TRANSFORM_PAK                           //GC_BOSS_TRANSFORM (Notify surrounding players that Boss start to transform)
	PACKET_GC_ERROR_MSG_PAK                                //GC_ERROR_MSG (general error message for all features in the game)
	PACKET_CG_CHAT_PAK                                     //CG_CHAT
	PACKET_GC_CHAT_PAK                                     //GC_CHAT
	PACKET_CG_RIDE_LEAVE_PAK                               //CG_RIDE_LEAVE
	PACKET_CG_RIDE_PAK                                     //CG_RIDE
	PACKET_GC_RIDE_PAK                                     //GC_RIDE
	PACKET_GC_ROBOT_CREATE_PAK                             //GC_ROBOT_CREATE
	PACKET_CG_LOGIN_ROBOT_PAK                              //Robot client ask login
	PACKET_GC_USER_PARAM_PAK                               //user param
	PACKET_CG_LEAVE_GAME_OK_PAK                            //recive leave game ok
	PACKET_GC_SYN_USER_INFO_PAK                            //
	PACKET_GC_BOX_NOTICE_PAK                               //GC_BOX_NOTICE
	PACKET_GC_BOX_LIST_PAK                                 //GC_BOX_LIST
	PACKET_CG_BOX_OPEN_PAK                                 //CG_BOX_OPEN
	PACKET_CG_ROBOT_ATTR_PAK                               //CG_ROBOT_ATTR
	PACKET_GC_PASSIVE_SKILL_LIST_PAK                       //
	PACKET_CG_NEARDEATH_KILL_PAK                           //CG_NEARDEATH_KILL
	PACKET_GC_NEARDEATH_KILL_PAK                           //GC_NEARDEATH_KILL
	PACKET_CG_NEARDEATH_SAVE_PAK                           //CG_NEARDEATH_SAVE
	PACKET_GC_NEARDEATH_SAVE_PAK                           //GC_NEARDEATH_SAVE
	PACKET_GC_TESTDELAY_PAK                                //GC_TESTDELAY
	PACKET_CG_TESTDELAY_PAK                                //CG_TESTDELAY
	PACKET_GC_SHOW_KICK_COUNTDWON_PAK                      //GC_SHOW_KICK_COUNTDWON
	PACKET_GC_HIDE_KICK_COUNTDWON_PAK                      //GC_HIDE_KICK_COUNTDWON
	PACKET_GC_AUTO_ADDHP_PAK                               //GC_AUTO_ADDHP
	PACKET_GC_GLOBAL_NOTIC_PAK                             //GC_GLOBAL_NOTIC
	PACKET_GC_RANK_MODIFY_NOTIC_PAK                        //GC_RANK_MODIFY_NOTIC
	PACKET_GC_ATTKER_PARAM_PAK                             //GC_ATTKER_PARAM
	PACKET_CG_FASHION_CHANGE_PAK                           //
	PACKET_GC_FASHION_CHANGE_PAK                           //
	PACKET_CG_ADJUSTATK_PAK                                //
	PACKET_GC_ADJUSTATK_PAK                                //
	PACKET_GC_BAGTIP_PAK                                   //
	PACKET_CG_BATCH_SALE_PAK                               //
	PACKET_CG_EXCHANGE_JJGOLD_PAK                          //
	PACKET_GC_EXCHANGE_JJGOLD_RESULT_PAK                   //
	PACKET_CG_QUERYJJGOLD_PAK                              //
	PACKET_GC_QUERYJJGOLD_PAK                              //
	PACKET_CG_JJ2DIAMOD_PAK                                //
	PACKET_GC_JJ2DIAMOD_RESULT_PAK                         //
	PACKET_GC_JJAWARD_RESULT_PAK                           //
	PACKET_CG_GOODS_INFO_PAK                               //CG_GOODS_INFO
	PACKET_GC_GOODS_INFO_PAK                               //GC_GOODS_INFO
	PACKET_CG_GOODS_BUY_PAK                                //CG_GOODS_BUY
	PACKET_GC_GOODS_BUY_PAK                                //GC_GOODS_BUY
	PACKET_CG_SHOP_REFRESH_PAK                             //shop refresh
	PACKET_GC_SHOP_REFRESH_RET_PAK                         //shop refresh ret
	PACKET_CG_REQEXINFO_PAK                                //CG_ReqExInfo
	PACKET_GC_REQEXINFO_PAK                                //GC_
	PACKET_CG_ADDMP_PAK                                    //CG_AddMp
	PACKET_CG_ADDHP_PAK                                    //CG_AddMp
	PACKET_CG_GENERALMESSAGETIPS_PAK                       //CG_GeneralMessageTips
	PACKET_GC_GENERALMESSAGETIPS_PAK                       //GC_GeneralMessageTips
	PACKET_CG_MSTSELECT_PAK                                //CG_GeneralMessageTips
	PACKET_CG_CLIENT_CONFIG_PAK                            //save client config
	PACKET_CG_USE_LOTTERYTICKET_PAK                        //save client config
	PACKET_GC_USE_LOTTERYTICKET_PAK                        //save client config
	PACKET_CG_LOTTERYTICKET_INFO_PAK                       //save client config
	PACKET_GC_LOTTERYTICKET_INFO_PAK                       //save client config
	PACKET_CG_LOTTERYTICKET_RECORD_PAK                     //save client config
	PACKET_GC_LOTTERYTICKET_RECORD_PAK                     //save client config
	PACKET_GC_SCENETASKBEGIN_PAK                           //
	PACKET_GC_SCENETASKINFOSYN_PAK                         //
	PACKET_GC_SCENETASKCOMPLETEINFO_PAK                    //
	PACKET_GC_SCENETASKPLAYERINFOSYN_PAK                   //
	PACKET_CG_SCENETASKRECORD_PAK                          //
	PACKET_GC_SCENETASKRECORD_PAK                          //
	PACKET_GC_SYNDWORDDATA_PAK                             //
	PACKET_GC_SYNMONSTERCOSTPER_PAK                        //
	PACKET_CG_ENHANCEEQUIP_PAK                             //
	PACKET_GC_ENHANCEEQUIP_PAK                             //
	PACKET_CG_OPENENERGY_PAK                               //
	PACKET_GC_OPENENERGY_PAK                               //
	PACKET_GC_EQUIPSUITINFO_PAK                            //
	PACKET_CG_ACTIVITY_INFO_PAK                            //
	PACKET_GC_ACTIVITY_INFO_PAK                            //
	PACKET_CG_ACTIVITY_REWARD_PAK                          //
	PACKET_GC_ACTIVITY_REWARD_PAK                          //
	PACKET_GC_STARTBOSSCHALLENGETASK_PAK                   //
	PACKET_GC_ENDBOSSCHALLENGETASK_PAK                     //
	PACKET_CG_REQUESTRECHARGELIST_PAK                      //
	PACKET_GC_REQUESTRECHARGELISTRET_PAK                   //
	PACKET_CG_REQUESTRECHARGEORDERID_PAK                   //
	PACKET_GC_REQUESTRECHARGEORDERIDRET_PAK                //
	PACKET_GC_RECHARGESUC_PAK                              //
	PACKET_CG_EXCHANGE_GOLDJJINFO_PAK                      //
	PACKET_GC_EXCHANGE_GOLDJJINFO_RESULT_PAK               //
	PACKET_CG_ENHANCEATTR_PAK                              //enhance attr
	PACKET_GC_ENHANCEATTR_RET_PAK                          //enhance attr
	PACKET_GC_COPY_SCENE_NOTIFY_PAK                        //copy scene notify
	PACKET_CG_COPY_SCENE_NOTIFY_RET_PAK                    //copy scene notify ret
	PACKET_GC_CHANGE_SCENE_PAK                             //Change Scene
	PACKET_CG_CHANGE_SCENE_RET_PAK                         //Change Scene Ret
	PACKET_CG_TICKET_LOTTERY_PAK                           //ticket lottery
	PACKET_GC_TICKET_LOTTERY_RET_PAK                       //ticket lottery ret
	PACKET_CG_TICKETLOTTERY_RECORD_PAK                     //get ticket lottery record
	PACKET_GC_TICKETLOTTERY_RECORD_RET_PAK                 //get ticket lottery record ret
	PACKET_CG_COMPOUND_PAK                                 //Compound
	PACKET_GC_COMPOUNDRESULT_PAK                           //Compound
	PACKET_CG_DAILYTASK_INFO_PAK                           //Get Daily task info
	PACKET_GC_DAILYTASK_INFO_RET_PAK                       //Get Daily task info
	PACKET_GC_SYN_DAILYTASK_PROGRESS_PAK                   //syn Daily task Progress
	PACKET_CG_DAILYTASK_ACCEPT_PAK                         //accept Daily task
	PACKET_GC_DAILYTASK_ACCEPT_RET_PAK                     //accept Daily task ret
	PACKET_CG_DAILYTASK_REFRESH_PAK                        //refresh Daily task
	PACKET_GC_DAILYTASK_REFRESH_RET_PAK                    //refresh Daily task
	PACKET_CG_DAILYTASK_RECORD_PAK                         //get Daily task  record
	PACKET_GC_DAILYTASK_RECORD_RET_PAK                     //get Daily task  record ret
	PACKET_GC_SYN_MODULELOCK_DATA_PAK                      //module lock syn info鍔熻兘瑙ｉ攣鏁版嵁鍚屾
	PACKET_CG_STORY_MISSION_INFO_PAK                       //get mission info
	PACKET_GC_STORY_MISSION_INFO_PAK                       // syn mission info
	PACKET_CG_GET_STORY_MISSION_REWARDS_PAK                // get mission reward
	PACKET_GC_GET_STORY_MISSION_REWARDS_RET_PAK            // get mission reward ret
	PACKET_GC_ITEM_NOTICE_LIST_PAK                         //GC_ITEM_NOTICE_LIST
	PACKET_GC_KICK_OUT_PAK                                 //Reconnection kick out
	PACKET_GC_CLIENT_CONFIG_PAK                            //GC_CLIENT_CONFIG
	PACKET_CG_REQUESTMONSTERSCORERECORD_PAK                //CG_RequestMonsterScore
	PACKET_GC_REQUESTMONSTERSCORERECORD_PAK                //GC_RequestMonsterScore
	PACKET_CG_REQUESTMONSTERSCOREINTERVAL_PAK              //CG_RequestMonsterScoreInterval
	PACKET_GC_REQUESTMONSTERSCOREINTERVAL_PAK              //GC_RequestMonsterScoreInterval
	PACKET_CG_CHANGE_TITLE_PAK                             //CG_ChangeTitle
	PACKET_GC_CHANGE_TITLE_RET_PAK                         //GC_ChangeTitle
	PACKET_GC_TITLE_LIST_PAK                               //GC_Title_LIST
	PACKET_GC_SYN_TITLE_APPEARANCE_PAK                     //GC_SYN_TITLE_APPEARANCE
	PACKET_GC_SYN_DEL_TITLE_PAK                            //GC_SYN_DEL_TITLE
	PACKET_GC_SYN_ADD_TITLE_PAK                            //GC_SYN_ADD_TITLE
	PACKET_CG_USE_CONSUM_ITEM_PAK                          //CG_USE_CONSUM_ITEM
	PACKET_GC_USE_CONSUM_ITEM_PAK                          //GC_USE_CONSUM_ITEM
	PACKET_GC_SYN_LIMITTIMES_PAK                           //GC_Syn_LimitTimes
	PACKET_CG_GET_FREE_MANA_PAK                            //CG_Get_Free_Mana
	PACKET_GC_GET_FREE_MANA_PAK                            //GC_Get_Free_Mana
	PACKET_CG_GET_MANA_GIF_BAG_INFO_PAK                    //CG_Get_Mana_Gif_Bag_Info
	PACKET_GC_GET_MANA_GIFT_BAG_PAK                        //GC_Get_Mana_Gift_Bag
	PACKET_GC_FUNCTIONPETLIST_PAK                          //
	PACKET_GC_FUNCTIONPET_NEW_PAK                          //
	PACKET_CG_FUNCTIONPETCONFIG_PAK                        //
	PACKET_GC_FUNCTIONPETCONFIG_PAK                        //
	PACKET_CG_SUMMONFUNCTIONPET_PAK                        //
	PACKET_GC_SUMMONFUNCTIONPET_RET_PAK                    //
	PACKET_CG_RESTFUNCTIONPET_PAK                          //
	PACKET_GC_RESTFUNCTIONPET_RET_PAK                      //
	PACKET_GC_FUNCTIONPET_EXPIRE_PAK                       //
	PACKET_GC_SYN_SYSTEMRECRODSTATE_PAK                    //
	PACKET_CG_GET_SYSTEMRECROD_PAK                         //
	PACKET_GC_GET_SYSTEMRECROD_PAK                         //
	PACKET_GC_SYNC_COMMONDATA_PAK                          //sys com data
	PACKET_CG_SIGNACTIVITY_INFO_PAK                        //Get sign activity info
	PACKET_GC_SIGNACTIVITY_INFO_RET_PAK                    //Get sign activity info
	PACKET_GC_SYN_SIGNACTIVITYTASK_PROGRESS_PAK            //syn sign activity task Progress
	PACKET_CG_ONEKEYCHANGEFASHION_PAK                      //
	PACKET_GC_ONEKEYCHANGEFASHIONRET_PAK                   //
	PACKET_CG_ONEKEYUPGRADEFASHION_PAK                     //
	PACKET_GC_ONEKEYUPGRADEFASHIONRET_PAK                  //
	PACKET_CG_ONEKEYFUSION_PAK                             //
	PACKET_GC_ONEKEYFUSIONRET_PAK                          //
	PACKET_CG_SINGLEEQUIPMENTUPGRADE_PAK                   //
	PACKET_GC_SINGLEEQUIPMENTUPGRADERET_PAK                //
	PACKET_CG_PLAYERSANDMONSTERSPOSREQ_PAK                 //
	PACKET_GC_PLAYERSANDMONSTERSPOSRET_PAK                 //

	PACKET_CG_MAX //消息类型的最大值
)
