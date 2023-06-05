// 编辑对话名称
const conversationName = document.getElementById("conversation-name");
const editButton = document.querySelector('.edit-button');
const saveButton = document.querySelector('.save-button');
const cancelButton = document.querySelector('.cancel-button');
editButton.addEventListener("click", () => {
    const nameInput = document.createElement("input");
    nameInput.type = "text";
    nameInput.value = conversationName.innerText;
    conversationName.innerHTML = "";
    conversationName.appendChild(nameInput);
    conversationName.appendChild(saveButton);
    conversationName.appendChild(cancelButton);
    conversationName.classList.add('editing');
    editButton.style.display = "none";
    saveButton.style.display = "inline"
    cancelButton.style.display = "inline"
});
saveButton.addEventListener("click", () => {
    const conversationId = document.querySelector(".active").dataset.conversationid;
    const nameInput = document.querySelector("#conversation-name input");
    fetch(`/chat/conversation/${conversationId}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `name=${nameInput.value}`
    }).then( response => {
        conversationName.classList.remove('editing');
        window.location.href = response.url
    });
});
cancelButton.addEventListener("click", () => {
    conversationName.classList.remove('editing');
    window.location.reload()
});

// 删除对话
const deleteButtons = document.querySelectorAll('.delete-button');
deleteButtons.forEach(button => {
    button.addEventListener('click', () => {
        const conversationId = button.parentNode.dataset.conversationid;
        fetch(`/chat/conversation/${conversationId}`, {
            method: "DELETE"
        }).then( response => {
            if (conversationId === document.querySelector(".active").dataset.conversationid) {
                window.location.href = response.url
            } else { window.location.reload() }
        });
    });
});

// 发送消息
const inputButton = document.querySelector('.input-area-button');
inputButton.addEventListener('click', () => {
    const content = document.querySelector('textarea[name="content"]').value
    const conversationId = document.querySelector(".active").dataset.conversationid;
    const send_message = {
        content: content,
        userID: 1, // 临时用户ID
        conversationID: conversationId
    }
    const received_message = {
        content: "正在输入...",
        userID: 0, // 系统固定ID
        conversationID: conversationId
    }
    inputButton.innerHTML = '<svg t="1683282636329" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11639" width="48" height="48"><path d="M524.8 748.8 320 896c-12.8 6.4 12.8 6.4 12.8 6.4l377.6 0c19.2 0 0-25.6 0-25.6-57.6-83.2-108.8-115.2-140.8-134.4C537.6 723.2 524.8 748.8 524.8 748.8z" p-id="11640" fill="#106CFF"></path><path d="M684.8 179.2 364.8 179.2c0 0-25.6 83.2 96 166.4 0 0 51.2 32 64 115.2l0 6.4c0 0 0 0 0 0 0 0 0 0 0 0L524.8 460.8c6.4-83.2 64-115.2 64-115.2C710.4 262.4 684.8 179.2 684.8 179.2z" p-id="11641" fill="#106CFF"></path><path d="M659.2 396.8c134.4-102.4 153.6-243.2 153.6-307.2l32 0c25.6 0 44.8-19.2 44.8-44.8l0 0c0-25.6-19.2-44.8-44.8-44.8L217.6 0C192 0 172.8 19.2 172.8 44.8l0 0c0 25.6 19.2 44.8 44.8 44.8l32 0c0 64 19.2 204.8 153.6 307.2C448 428.8 473.6 467.2 473.6 512c0 70.4-64 121.6-70.4 121.6-12.8 6.4-153.6 108.8-153.6 300.8L217.6 934.4c-25.6 0-44.8 19.2-44.8 44.8l0 0C172.8 1004.8 192 1024 217.6 1024l627.2 0c25.6 0 44.8-19.2 44.8-44.8l0 0c0-25.6-19.2-44.8-44.8-44.8l-32 0c0-192-147.2-294.4-153.6-300.8C652.8 627.2 588.8 576 588.8 512 588.8 467.2 608 428.8 659.2 396.8zM556.8 512c0 83.2 76.8 140.8 83.2 147.2l0 0c0 0 140.8 96 140.8 275.2l-512 0c0-179.2 140.8-275.2 140.8-275.2l0 0c6.4 0 83.2-64 83.2-147.2 0-51.2-25.6-96-83.2-134.4C294.4 281.6 275.2 147.2 275.2 89.6l505.6 0c0 57.6-19.2 185.6-140.8 281.6C588.8 416 556.8 460.8 556.8 512z" p-id="11642" fill="#106CFF"></path></svg>'
    inputButton.setAttribute('disabled', true)
    let inputArea = document.getElementsByClassName('input-area')[0]
    inputArea.getElementsByTagName('textarea')[0].setAttribute('disabled', true)
    appendMessageToChatArea(send_message)
    appendMessageToChatArea(received_message)
    fetch(`/chat/conversation/${conversationId}/message`, {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `content=${content}`
    }).then( response => {
        window.location.href = response.url
    });
});
const textarea = document.querySelector('textarea');
textarea.addEventListener('keydown', event => {
    if (event.key === 'Enter' && !event.shiftKey) {
        event.preventDefault();
        inputButton.click();
    }
});

// 清空消息
const clearButton = document.querySelector('.clear-button');
clearButton.addEventListener('click', () => {
    const conversationId = document.querySelector(".active").dataset.conversationid;
    fetch(`/chat/conversation/${conversationId}/message`, {
        method: "DELETE"
    }).then( response => {
        window.location.href = response.url
    });
});

// 切换对话
function selectConversation(conversationID) {
    // 遍历所有li元素，将选中的元素设置为active，其他元素移除active
    document.querySelectorAll('.chat-list li').forEach(function(li) {
        if (li.getAttribute('data-conversationID') === conversationID) {
            li.classList.add('active');
        } else {
            li.classList.remove('active');
        }
    });
    // 切换页面到选中的对话框
    window.location.href = '/chat/conversation/' + conversationID;
}

// 折叠对话列表
function toggleLeftPanel() {
    const container = document.querySelector('.container');
    container.classList.toggle('collapsed');
}

// 将消息添加到聊天区域
function appendMessageToChatArea(message) {
    let chatArea = document.getElementsByClassName('chat-area')[0]
    let li = document.createElement('li')
    li.setAttribute('class', message.userID === 0 ? 'received' : 'send')
    li.innerHTML = `
      <div class="avatar">
        ${message.userID === 0 ? '<img src="/img/ai-chat.png" alt="">' : '<img src="/img/github.svg" alt="">'}
      </div>
      <div class="message">${message.content}</div>
      <div class="time">
      ${message.userID === 0 ? '<span>现在</span>' : '<span>刚刚</span>'}
      </div>
    `
    chatArea.appendChild(li)
}