select *
  from (select distinct t3.id,
                        t3.url_type    as urlType,
                        t3.app_id      as appId,
                        t1.id          as plateId,
                        t1.name        as plateName,
                        t1.source      as plateSource,
                        t1.source_id   as plateSourceId,
                        t1.source_name as plateSourceName,
                        t1.icon_choose as plateIconChoose,
                        t1.icon        as plateIcon,
                        t1.index       as plateIndex,
                        t1.status      as plateStatus,
                        t2.id          as moduleId,
                        t2.name        as moduleName,
                        t2.source      as moduleSource,
                        t2.source_id   as moduleSourceId,
                        t2.source_name as moduleSourceName,
                        t2.layout      as moduleLayout,
                        t2.display_num as moduleDisplayNum,
                        t2.frequency   as moduleFrequency,
                        t2.index       as moduleIndex,
                        t2.status      as moduleStatus
          from (select id,name,source,source_id,source_name from t_plate where id > 0 )  t1
              join (select id,name,source,source_id,source_name from t_module where id > 0 ) t2
            on t1.id = t2.plate_id
          left join t_module_content t3
            on t2.id = t3.module_id 
         where t1.status = 1
         order by "t1"."id",
                  "t1"."index",
                  "t2"."id",
                  "t2"."index",
                  "t3"."index") t
 where id in (select id from t_plate where status = 1)
   and name = ''