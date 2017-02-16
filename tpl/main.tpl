{{template "header"}}
{{define "main"}}

  <div class="main">
    <div class="folders">
      <!-- Folder List -->
        <script language="javascript">
            $.get("/folders").then(function(data) {
                    $(".folders").html(data);
                }
            );
        </script>
      <!-- /Folder List -->
    </div>
    <div class="message_box">
      <!-- Message List -->
        <div class="message_list">
        {{template "messages"}}
        </div>
      <!-- /Message List -->

      <!-- Message -->
        <div class="message">
        {{template "message"}}
        </div>
      <!-- /Message -->

    </div>
  </div>
{{end}}

{{template "footer"}}
