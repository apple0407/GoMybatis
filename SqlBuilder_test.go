package GoMybatis

import (
	"testing"
	"fmt"
	"reflect"
	"github.com/zhuxiujia/GoMybatis/utils"
	"time"
)

//测试sql生成tps
func Test_SqlBuilder_Tps(t *testing.T) {
	var mapper = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper>
    <!--List<Activity> selectByCondition(@Param("name") String name,@Param("startTime") Date startTime,@Param("endTime") Date endTime,@Param("index") Integer index,@Param("size") Integer size);-->
    <!-- 后台查询产品 -->
    <select id="selectByCondition">
        select * from biz_activity where delete_flag=1
        <if test="name != ''">
            and name like concat('%',#{name},'%')
        </if>
        <if test="startTime != ''">
            and create_time >= #{startTime}
        </if>
        <if test="endTime != ''">
            and create_time &lt;= #{endTime}
        </if>
        order by create_time desc
        <if test="page >= 0 and size != 0">limit #{page}, #{size}</if>
    </select>
    <!--int countByCondition(@Param("name")String name,@Param("startTime") Date startTime, @Param("endTime")Date endTime);-->
    <select id="countByCondition">
        select count(id) from biz_activity where delete_flag=1
        <if test="name != ''">
            and name like concat('%',#{name},'%')
        </if>
        <if test="startTime != ''">
            and create_time >= #{startTime}
        </if>
        <if test="endTime != ''">
            and create_time &lt;= #{endTime}
        </if>
    </select>
    <!--List<Activity> selectAll();-->
    <select id="selectAll">
        select * from biz_activity where delete_flag=1 order by create_time desc
    </select>
    <!--Activity selectByUUID(@Param("uuid")String uuid);-->
    <select id="selectByUUID">
        select * from biz_activity
        where uuid = #{uuid}
        and delete_flag = 1
    </select>
    <select id="selectById">
        select * from biz_activity
        where id = #{id}
        and delete_flag = 1
    </select>
    <select id="selectByIds">
        select * from biz_activity
        where delete_flag = 1
        and id in
        <foreach collection="ids" item="item" index="index" open="(" close=")">
            #{item}
        </foreach>
    </select>
    <update id="deleteById">
        update biz_activity
        set delete_flag = 0
        where id = #{id}
    </update>
    <update id="updateById">
        update biz_activity
        <set>
            <if test="name != ''">name = #{name,jdbcType=VARCHAR},</if>
            <if test="pcLink != ''">pc_link = #{pcLink,jdbcType=VARCHAR},</if>
            <if test="h5Link != ''">h5_link = #{h5Link,jdbcType=VARCHAR},</if>
            <if test="remark != ''">remark = #{remark,jdbcType=VARCHAR},</if>
            <if test="createTime != ''">create_time = #{createTime,jdbcType=TIMESTAMP},</if>
            <if test="deleteFlag != ''">delete_flag = #{deleteFlag},</if>
        </set>
        where id = #{id} and delete_flag = 1
    </update>
    <insert id="insert">
        insert into biz_activity
        <trim prefix="(" suffix=")" suffixOverrides=",">
            <if test="id != ''">id,</if>
            <if test="name != ''">name,</if>
            <if test="pcLink != ''">pc_link,</if>
            <if test="h5Link != ''">h5_link,</if>
            <if test="remark != ''">remark,</if>
            <if test="createTime != ''">create_time,</if>
            <if test="deleteFlag != ''">delete_flag,</if>
        </trim>

        <trim prefix="values (" suffix=")" suffixOverrides=",">
            <if test="id != ''">#{id,jdbcType=VARCHAR},</if>
            <if test="name != ''">#{name,jdbcType=VARCHAR},</if>
            <if test="pcLink != ''">#{pcLink,jdbcType=VARCHAR},</if>
            <if test="h5Link != ''">#{h5Link,jdbcType=VARCHAR},</if>
            <if test="remark != ''">#{remark,jdbcType=VARCHAR},</if>
            <if test="createTime != ''">#{createTime,jdbcType=TIMESTAMP},</if>
            <if test="deleteFlag != ''">#{deleteFlag},</if>
        </trim>
    </insert>
</mapper>`
	var mapperTree = LoadMapperXml([]byte(mapper))
	var builder = GoMybatisSqlBuilder{}.New(GoMybatisExpressionTypeConvert{}, GoMybatisSqlArgTypeConvert{})
	var paramMap = make(map[string]SqlArg)
	paramMap["name"] = SqlArg{
		Value: "ss",
		Type:  reflect.TypeOf("ss"),
	}
	defer utils.CountMethodTps(100000, time.Now(), "Test_SqlBuilder")
	for i := 0; i < 100000; i++ {
		//var sql, e =
		builder.BuildSql(paramMap, mapperTree[0])
		//fmt.Println(sql, e)
	}
	fmt.Println("done")
}
