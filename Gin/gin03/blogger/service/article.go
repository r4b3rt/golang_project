package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get article list failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	// 获取分类ids
	categoryIds := getCategoryIds(articleInfoList)
	// 依据分类id 获取分类信息
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId //分类id
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids { // 去重
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// 根据分类id获取文章信息 (用于文章上下页)
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		fmt.Printf("get artcle list failed,err %v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	categoryIds := getCategoryIds(articleInfoList)       // 过滤到重复的分类id
	categoryList, err := db.GetCategoryList(categoryIds) //  根据分类id 获取分类信息
	if err != nil {
		fmt.Printf("2 get category list failed, err:%v\n", err)
		return
	}

	for _, article := range articleInfoList {
		fmt.Printf("content:%s\n", article.Summary)
		// 初始化分页结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}
